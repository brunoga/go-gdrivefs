package filesystem

import (
	"github.com/hanwen/go-fuse/fuse/nodefs"
)

// file is a file that does nothing but log all operations called on it.
type file struct {
	inode *nodefs.Inode
	root  *rootNode

	*loggingFile
}

// newFile returns a new file instance.
func newFile(rootNode *rootNode) nodefs.File {
	f := &file{
		nil,
		rootNode,
		newLoggingFile(true).(*loggingFile),
	}
	f.setLogPrefix("BaseFile")

	return f
}

// nodefs.File generic interface methods.

func (f *file) SetInode(inode *nodefs.Inode) {
	f.loggingFile.SetInode(inode)

	f.inode = inode
}
