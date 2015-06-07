package filesystem

import (
	"github.com/hanwen/go-fuse/fuse/nodefs"
)

// baseFile is a file that does nothing but log all operations called on it.
type baseFile struct {
	inode *nodefs.Inode

	*loggingFile
}

// NewBaseFile returns a new baseFile instance.
func NewBaseFile(shouldLog bool) nodefs.File {
	return &baseFile{
		nil,
		NewLoggingFile(true).(*loggingFile),
	}
}

// nodefs.File generic interface methods.

func (f *baseFile) SetInode(inode *nodefs.Inode) {
	f.loggingFile.SetInode(inode)

	f.inode = inode
}
