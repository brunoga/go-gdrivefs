package filesystem

import (
	"fmt"

	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/nodefs"

	drive "google.golang.org/api/drive/v2"
)

// directoryNode is a node that contains other nodes (i.e. a directory). It only
// implements methods related to directory manipulation, deferring all other
// calls to baseNode.
//
// TODO(bga): Implement relevant methods.
type directoryNode struct {
	driveFiles []*drive.File

	*baseNode
}

// NewDirectoryNode returns a new directoryNode instance.
func NewDirectoryNode() nodefs.Node {
	n := &directoryNode{
		nil,
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
	driveFile, err := n.getRootNode().gdriveHandler.GetFileByName(name, id)
	if err != nil {
		// Could not retrieve data for file.
		n.log(fmt.Sprintf("Error retrieving file data : %s", err))

		return nil, fuse.EIO
	}

	if driveFile == nil {
		// No file found.
		return nil, fuse.ENOENT
	}

	isDir := fillAttr(n.loggingNode, driveFile, out)

	var newNode nodefs.Node
	if isDir {
		// Setup directory node.
		newNode = NewDirectoryNode()
		newNode.(*directoryNode).driveEntry = driveFile
		newNode.(*directoryNode).setRootNode(n.getRootNode())
	} else {
		// Setup file node.
		newNode = NewFileNode()
		newNode.(*fileNode).driveEntry = driveFile
		newNode.(*fileNode).setRootNode(n.getRootNode())
	}

	// Allocate a new Inode.
	newInode := n.Inode().NewChild(name, isDir, newNode)

	return newInode, fuse.OK
}

func (n *directoryNode) OpenDir(
	context *fuse.Context) ([]fuse.DirEntry, fuse.Status) {
	n.baseNode.OpenDir(context)

	if n.driveFiles == nil {
		driveFiles, err := n.getRootNode().gdriveHandler.GetFileList(
			n.driveEntry.Id)
		if err != nil {
			n.log(fmt.Sprintf("Error retrieving file list : %s",
				err))
			return nil, fuse.EIO
		}

		n.driveFiles = driveFiles
	}

	dirEntries := make([]fuse.DirEntry, 0, len(n.driveFiles))
	for _, driveFile := range n.driveFiles {
		if driveFile.Labels.Trashed {
			// Ignore items in the trash.
			continue
		}

		var mode uint32
		if driveFile.MimeType == "application/vnd.google-apps.folder" {
			mode = fuse.S_IFDIR | 0755
		} else {
			mode = fuse.S_IFREG | 0644
		}

		dirEntries = append(dirEntries, fuse.DirEntry{
			Name: driveFile.Title,
			Mode: mode,
		})
	}

	return dirEntries, fuse.OK
}
