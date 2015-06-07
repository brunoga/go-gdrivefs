package filesystem

import (
	"log"
	"time"

	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/nodefs"
)

// loggingFile is a file that does nothing but log all operations called on it.
type loggingFile struct {
	shouldLog bool
	logPrefix string
}

// NewLoggingFile returns a new loggingFile instance. The shouldLog parameter
// determines if it should log operations called on it or not.
func NewLoggingFile(shouldLog bool) nodefs.File {
	return &loggingFile{
		shouldLog,
		"LoggingFile",
	}
}

// log logs the given text using the default logger.
func (f *loggingFile) log(text string) {
	if f.shouldLog {
		log.Printf("%s : %p : %s", f.logPrefix, f, text)
	}
}

// setLogPrefix sets the log text prefix to logPrefix.
func (f *loggingFile) setLogPrefix(logPrefix string) {
	f.logPrefix = logPrefix
}

// nodefs.File interface methods.

func (f *loggingFile) SetInode(*nodefs.Inode) {
	f.log("SetInode()")
}

func (f *loggingFile) InnerFile() nodefs.File {
	f.log("InnerFile()")
	return nil
}

func (f *loggingFile) String() string {
	f.log("STring()()")
	return "loggingFile"
}

func (f *loggingFile) Read(buf []byte, off int64) (fuse.ReadResult, fuse.Status) {
	f.log("Read()")
	return nil, fuse.ENOSYS
}

func (f *loggingFile) Write(data []byte, off int64) (uint32, fuse.Status) {
	f.log("Write()")
	return 0, fuse.ENOSYS
}

func (f *loggingFile) Flush() fuse.Status {
	f.log("Flush()")
	return fuse.OK
}

func (f *loggingFile) Release() {
	f.log("Release()")
}

func (f *loggingFile) GetAttr(*fuse.Attr) fuse.Status {
	f.log("GetAttr()")
	return fuse.ENOSYS
}

func (f *loggingFile) Fsync(flags int) (code fuse.Status) {
	f.log("Fsync()")
	return fuse.ENOSYS
}

func (f *loggingFile) Utimens(atime *time.Time, mtime *time.Time) fuse.Status {
	f.log("Utimens()")
	return fuse.ENOSYS
}

func (f *loggingFile) Truncate(size uint64) fuse.Status {
	f.log("Truncate()")
	return fuse.ENOSYS
}

func (f *loggingFile) Chown(uid uint32, gid uint32) fuse.Status {
	f.log("Chown()")
	return fuse.ENOSYS
}

func (f *loggingFile) Chmod(perms uint32) fuse.Status {
	f.log("Chmod()")
	return fuse.ENOSYS
}

func (f *loggingFile) Allocate(off uint64, size uint64, mode uint32) (code fuse.Status) {
	f.log("Allocate()")
	return fuse.ENOSYS
}
