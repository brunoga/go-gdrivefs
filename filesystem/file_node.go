package filesystem

import (
	"os"

	"github.com/hanwen/go-fuse/fuse"
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

// nodefs.Node file-related interface methods.

func (n *fileNode) Open(flags uint32,
	context *fuse.Context) (file nodefs.File, code fuse.Status) {
	// Currrently we have a read-only file system.
	if (flags & (uint32(os.O_WRONLY | os.O_RDWR | os.O_TRUNC))) != 0 {
		return nil, fuse.EROFS
	}

	// Currently we do not return a proper file. This is ok and we can
	// handle everything at the node level.
	//
	// TODO(bga): Consider creating a proper file representation.
	return nil, fuse.OK
}
