package filesystem

import (
	"github.com/hanwen/go-fuse/fuse/nodefs"
)

// baseFile is a file that does nothing but log all operations called on it.
type baseFile struct {
	inode *nodefs.Inode
	root  *rootNode

	*loggingFile
}

// NewBaseFile returns a new baseFile instance.
func NewBaseFile(rootNode *rootNode) nodefs.File {
	f := &baseFile{
		nil,
		rootNode,
		NewLoggingFile(true).(*loggingFile),
	}
	f.setLogPrefix("BaseFile")

	return f
}

// nodefs.File generic interface methods.

func (f *baseFile) SetInode(inode *nodefs.Inode) {
	f.loggingFile.SetInode(inode)

	f.inode = inode
}
