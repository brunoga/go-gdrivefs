package filesystem

import (
	"github.com/hanwen/go-fuse/fuse/nodefs"
)

// fileNode is Node that represents a file. it only implements methods related
// to file manipulation, deferring all other calls to baseNode.
//
// TODO(bga): Implement relevant methods.
type fileNode struct {
	*baseNode
}

// NewFileNode returns a new fileNode instance.
func NewFileNode() nodefs.Node {
	n := &fileNode{
		NewBaseNode().(*baseNode),
	}
	n.setLogPrefix("FileNode")

	return n
}
