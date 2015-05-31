package filesystem

import (
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
