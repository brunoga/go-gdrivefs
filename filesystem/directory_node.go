package filesystem

import (
	"fmt"
	"time"

	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/nodefs"
)

// directoryNode is a node that contains other nodes (i.e. a directory). It only
// implements methods related to directory manipulation, deferring all other
// calls to baseNode.
//
// TODO(bga): Implement relevant methods.
type directoryNode struct {
	*baseNode
}

// NewDirectoryNode returns a new directoryNode instance.
func NewDirectoryNode() nodefs.Node {
	n := &directoryNode{
		NewBaseNode().(*baseNode),
	}
	n.setLogPrefix("DirectoryNode")

	return n
}

// nodefs.Node directory-related interface methods.

func (n *directoryNode) Lookup(out *fuse.Attr, name string,
	context *fuse.Context) (node *nodefs.Inode, code fuse.Status) {
	n.baseNode.Lookup(out, name, context)

	// Figure out our own directory id.
	var id string
	if n.driveEntry == nil {
		// We have no information about ourselves, so we must be the
		// root node.
		id = "root"
	} else {
		// We have information about ourselves. Use it.
		id = n.driveEntry.Id
	}

	// Call handler to obtain information about the required file in this
	// directory.
	c := n.getRootNode().gdriveHandler.GetFileByName(name, id)
	result := <-c

	items, err := result.Get()
	if err != nil {
		// Could not retrieve data for file.
		n.log(fmt.Sprintf("Error retrieving file data : %s", err))

		return nil, fuse.EIO
	}

	if len(items) == 0 {
		// No file found.
		return nil, fuse.ENOENT
	}

	n.log("HERE")

	item := items[0]

	// Set file size.
	out.Size = uint64(item.FileSize)

	// Set crested and modified dates.
	createdDate, err := time.Parse(time.RFC3339, item.CreatedDate)
	if err == nil {
		out.Ctime = uint64(createdDate.Unix())
	} else {
		n.log(fmt.Sprintf("Error parsing date : %s", err))
	}

	modifiedDate, err := time.Parse(time.RFC3339, item.ModifiedDate)
	if err == nil {
		out.Mtime = uint64(modifiedDate.Unix())
	} else {
		n.log(fmt.Sprintf("Error parsing date : %s", err))
	}

	// Set up file/directory entry.
	var isDir bool
	var newNode nodefs.Node
	if item.MimeType == "application/vnd.google-apps.folder" {
		// This is a directory.
		out.Mode = fuse.S_IFDIR | 0755

		newNode = NewDirectoryNode()

		isDir = true
	} else {
		// This is a file.
		out.Mode = fuse.S_IFREG | 0644

		newNode = NewFileNode()

		isDir = false
	}

	// Set data and rootnode for entry.
	newNode.(*baseNode).driveEntry = item
	newNode.(*baseNode).setRootNode(n.getRootNode())

	// Allocatea new Inode.
	newInode := n.Inode().NewChild(name, isDir, newNode)

	return newInode, fuse.OK
}
