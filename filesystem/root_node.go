package filesystem

import (
	"github.com/hanwen/go-fuse/fuse/nodefs"
)

// rootNode is the Node at the root of the filesystem. It implements filesystem
// specific methods and delegates all other calls to directoryNode (a rootNode
// is also a directoryNode).
//
// TODO(bga): Actualy implement the needed methods.
type rootNode struct {
	*directoryNode
}

// NewRootNode returns a new rootNode instance.
func NewRootNode() nodefs.Node {
	n := &rootNode{
		NewDirectoryNode().(*directoryNode),
	}
	n.setLogPrefix("RootNode")

	return n
}
