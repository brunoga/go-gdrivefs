package filesystem

import (
	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/nodefs"

	drive "google.golang.org/api/drive/v2"
)

// file is a file that does nothing but log all operations called on it.
type file struct {
	inode     *nodefs.Inode
	root      *rootNode
	driveFile *drive.File

	*loggingFile
}

// newFile returns a new file instance.
func newFile(rootNode *rootNode, driveFile *drive.File) nodefs.File {
	f := &file{
		nil,
		rootNode,
		driveFile,
		newLoggingFile(true).(*loggingFile),
	}
	f.setLogPrefix("BaseFile")

	return f
}

// nodefs.File interface methods.

func (f *file) SetInode(inode *nodefs.Inode) {
	f.loggingFile.SetInode(inode)

	f.inode = inode
}

func (f *file) Read(buf []byte, off int64) (fuse.ReadResult, fuse.Status) {
	h := f.root.gdriveHandler

	return h.ReadFile(f.driveFile, buf, off)
}
