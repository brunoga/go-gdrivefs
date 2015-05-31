package filesystem

import (
	"github.com/hanwen/go-fuse/fuse/nodefs"
)

// directoryNode is a node that contains other nodes (i.e. a directory). It only
// implements methods related to directory manipulation, deferring all other
// calls to loggingNode.
//
// TODO(bga): Actualy implement the needed methods.
type directoryNode struct {
	*loggingNode
}

// NewDirectoryNode returns a new directoryNode instance.
func NewDirectoryNode() nodefs.Node {
	n := &directoryNode{
		NewLoggingNode(true).(*loggingNode),
	}
	n.setLogPrefix("DirectoryNode")

	return n
}
