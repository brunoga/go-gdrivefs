package filesystem

import (
	"fmt"

	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/nodefs"

	drive "google.golang.org/api/drive/v2"
)

// baseNode implements the Inode-related methods in nodefs.Node. It delegates
// all other calls to loggingNode
type baseNode struct {
	inode *nodefs.Inode
	root  *rootNode

	driveEntry *drive.File

	*loggingNode
}

// newBaseNode returns a new baseNode instance.
func newBaseNode() nodefs.Node {
	n := &baseNode{
		nil,
		nil,
		nil,
		newLoggingNode(true).(*loggingNode),
	}
	n.setLogPrefix("BaseNode")

	return n
}

func (n *baseNode) getRootNode() *rootNode {
	return n.root
}

func (n *baseNode) setRootNode(root *rootNode) {
	n.root = root
}

// nodefs.Node generic interface methods.

func (n *baseNode) SetInode(node *nodefs.Inode) {
	n.loggingNode.SetInode(node)

	n.inode = node
}

func (n *baseNode) Inode() *nodefs.Inode {
	n.loggingNode.Inode()

	return n.inode
}

func (n *baseNode) GetAttr(out *fuse.Attr, file nodefs.File,
	context *fuse.Context) (code fuse.Status) {
	n.loggingNode.GetAttr(out, file, context)

	if n.driveEntry == nil {
		driveFile, err := n.getRootNode().gdriveHandler.GetFileById(
			"root")
		if err != nil {
			n.log(fmt.Sprintf(
				"Error retrieving file data : %s", err))
			return fuse.EIO
		}

		if driveFile == nil {
			return fuse.ENOENT
		}

		n.driveEntry = driveFile
	}

	fillAttr(n.loggingNode, n.driveEntry, out)

	return fuse.OK
}
