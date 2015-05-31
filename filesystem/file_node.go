package filesystem

import (
	"github.com/hanwen/go-fuse/fuse/nodefs"
)

// fileNode is Node that represents a file. it only implements methods related
// to file manipulation, deferring all other calls to loggingNode.
type fileNode struct {
	loggingNode
}

// NewFileNode returns a new fileNode instance.
func NewFileNode() nodefs.Node {
	n := &fileNode{}
	n.setLogPrefix("FileNode")

	return n
}
