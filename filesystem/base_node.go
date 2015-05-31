package filesystem

import (
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

// NewBaseNode returns a new baseNode instance.
func NewBaseNode() nodefs.Node {
	n := &baseNode{
		nil,
		nil,
		nil,
		NewLoggingNode(true).(*loggingNode),
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

// nodefs.Node Inode-related interface methods.

func (n *baseNode) SetInode(node *nodefs.Inode) {
	n.loggingNode.SetInode(node)

	n.inode = node
}

func (n *baseNode) Inode() *nodefs.Inode {
	n.loggingNode.Inode()

	return n.inode
}
