package filesystem

import (
	"github.com/brunoga/go-gdrivefs/gdrive"
	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/nodefs"
)

// rootNode is the Node at the root of the filesystem. It implements filesystem
// specific methods and delegates all other calls to directoryNode (a rootNode
// is also a directoryNode).
//
// TODO(bga): Implement relevant methods.
type rootNode struct {
	gdriveHandler *gdrive.Handler

	*directoryNode
}

// NewRootNode returns a new rootNode instance.
func NewRootNode(gdriveHandler *gdrive.Handler) nodefs.Node {
	n := &rootNode{
		gdriveHandler,
		NewDirectoryNode().(*directoryNode),
	}
	n.setLogPrefix("RootNode")
	n.setRootNode(n)

	return n
}

// nodefs.Node root-related interface methods.

func (n *rootNode) OnMount(conn *nodefs.FileSystemConnector) {
	// TODO(bga): Add any filesystem initializetion code here.
	n.directoryNode.OnMount(conn)
}

func (n *rootNode) OnUnmount() {
	// TODO(bga): Add any filesystem teardown code here.
	n.directoryNode.OnUnmount()
}

func (n *rootNode) StatFs() *fuse.StatfsOut {
	// TODO(bga): Add code to fill the StatfsOut structure here.
	return n.directoryNode.StatFs()
}
