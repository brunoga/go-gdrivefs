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

// newFileNode returns a new fileNode instance.
func newFileNode() nodefs.Node {
	n := &fileNode{
		newBaseNode().(*baseNode),
	}
	n.setLogPrefix("FileNode")

	return n
}

// nodefs.Node file-related interface methods.

func (n *fileNode) Open(flags uint32,
	context *fuse.Context) (file nodefs.File, code fuse.Status) {
	// TODO(bga): This method should most likelly move to base_node.go as
	// Open() can be called on directories. OTOH, it might be better to have
	// a separate Open() implementation in this case. It will all depend on
	// the amount of code that could be resused.

	n.baseNode.Open(flags, context)

	// Currrently we have a read-only file system.
	if (flags & (uint32(os.O_WRONLY | os.O_RDWR | os.O_TRUNC))) != 0 {
		return nil, fuse.EROFS
	}

	newFile := NewBaseFile(n.getRootNode())

	// Currently we do not return a proper file. This is ok and we can
	// handle everything at the node level.
	//
	// TODO(bga): Consider creating a proper file representation.
	return newFile, fuse.OK
}

func (n *fileNode) Read(file nodefs.File, dest []byte, off int64,
	context *fuse.Context) (fuse.ReadResult, fuse.Status) {
	n.baseNode.Read(file, dest, off, context)

	if file != nil {
		// Delegate to file.
		return file.Read(dest, off)
	}

	// We don't have a file and we also do not implement reading at the
	// node level.
	return nil, fuse.ENOSYS
}

func (n *fileNode) Write(file nodefs.File, data []byte, off int64,
	context *fuse.Context) (written uint32, code fuse.Status) {
	n.baseNode.Write(file, data, off, context)

	if file != nil {
		// Delegate to file.
		return file.Write(data, off)
	}

	// We don't have a file and we also do not implement writting at the
	// node level.
	return 0, fuse.ENOSYS
}
