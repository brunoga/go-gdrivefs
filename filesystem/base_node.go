package filesystem

import (
	"github.com/hanwen/go-fuse/fuse/nodefs"
)

// baseNode implements the Inode-related methods in nodefs.Node. It delegates
// all other calls to loggingNode
type baseNode struct {
	inode *nodefs.Inode

	*loggingNode
}

// NewBaseNode returns a new baseNode instance.
func NewBaseNode() nodefs.Node {
	n := &baseNode{
		nil,
		NewLoggingNode(true).(*loggingNode),
	}
	n.setLogPrefix("BaseNode")

	return n
}

// nodefs.Node Inode menipulation interface methods.

func (n *baseNode) SetInode(node *nodefs.Inode) {
	n.loggingNode.SetInode(node)

	n.inode = node
}

func (n *baseNode) Inode() *nodefs.Inode {
	n.loggingNode.Inode()

	return n.inode
}
