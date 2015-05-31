package filesystem

import (
	"log"
	"time"

	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/nodefs"
)

// loggingNode is a node that does nothing but log all operations called on it.
type loggingNode struct {
	shouldLog bool
	logPrefix string
}

// NewLoggingNode returns a new loggingNodei instance. The shouldLog parameter
// determines if it should log operations called on it or not.
func NewLoggingNode(shouldLog bool) nodefs.Node {
	return &loggingNode{
		shouldLog,
		"LoggingNode",
	}
}

// log logs the given text using the default logger.
func (n *loggingNode) log(text string) {
	if n.shouldLog {
		log.Printf("%s : %p : %s", n.logPrefix, n, text)
	}
}

// setLogPrefix sets the log text prefix to logPrefix.
func (n *loggingNode) setLogPrefix(logPrefix string) {
	n.logPrefix = logPrefix
}

// nodefs.Node interface methods.

func (n *loggingNode) OnUnmount() {
	n.log("Unmount()")
}

func (n *loggingNode) OnMount(conn *nodefs.FileSystemConnector) {
	n.log("OnMount()")
}

func (n *loggingNode) StatFs() *fuse.StatfsOut {
	n.log("StatFs()")
	return nil
}

func (n *loggingNode) SetInode(node *nodefs.Inode) {
	n.log("SetInode()")
}

func (n *loggingNode) Deletable() bool {
	n.log("Deletable()")
	return true
}

func (n *loggingNode) Inode() *nodefs.Inode {
	n.log("Inode()")
	return nil
}

func (n *loggingNode) OnForget() {
	n.log("OnForget()")
}

func (n *loggingNode) Lookup(out *fuse.Attr, name string,
	context *fuse.Context) (node *nodefs.Inode, code fuse.Status) {
	n.log("Lookup()")
	return nil, fuse.ENOSYS
}

func (n *loggingNode) Access(mode uint32,
	context *fuse.Context) (code fuse.Status) {
	n.log("Access()")
	return fuse.ENOSYS
}

func (n *loggingNode) Readlink(c *fuse.Context) ([]byte, fuse.Status) {
	n.log("Readlink()")
	return nil, fuse.ENOSYS
}

func (n *loggingNode) Mknod(name string, mode uint32, dev uint32,
	context *fuse.Context) (newNode *nodefs.Inode, code fuse.Status) {
	n.log("Mknod()")
	return nil, fuse.ENOSYS
}
func (n *loggingNode) Mkdir(name string, mode uint32,
	context *fuse.Context) (newNode *nodefs.Inode, code fuse.Status) {
	n.log("Mkdir()")
	return nil, fuse.ENOSYS
}
func (n *loggingNode) Unlink(name string,
	context *fuse.Context) (code fuse.Status) {
	n.log("Unlink()")
	return fuse.ENOSYS
}
func (n *loggingNode) Rmdir(name string,
	context *fuse.Context) (code fuse.Status) {
	n.log("Rmdir()")
	return fuse.ENOSYS
}
func (n *loggingNode) Symlink(name string, content string,
	context *fuse.Context) (newNode *nodefs.Inode, code fuse.Status) {
	n.log("Symlink()")
	return nil, fuse.ENOSYS
}

func (n *loggingNode) Rename(oldName string, newParent nodefs.Node,
	newName string, context *fuse.Context) (code fuse.Status) {
	n.log("Rename()")
	return fuse.ENOSYS
}

func (n *loggingNode) Link(name string, existing nodefs.Node,
	context *fuse.Context) (newNode *nodefs.Inode, code fuse.Status) {
	n.log("Link()")
	return nil, fuse.ENOSYS
}

func (n *loggingNode) Create(name string, flags uint32, mode uint32,
	context *fuse.Context) (file nodefs.File, newNode *nodefs.Inode,
	code fuse.Status) {
	n.log("Create()")
	return nil, nil, fuse.ENOSYS
}

func (n *loggingNode) Open(flags uint32,
	context *fuse.Context) (file nodefs.File, code fuse.Status) {
	n.log("Open()")
	return nil, fuse.ENOSYS
}

func (n *loggingNode) Flush(file nodefs.File, openFlags uint32,
	context *fuse.Context) (code fuse.Status) {
	n.log("Flush()")
	return fuse.ENOSYS
}

func (n *loggingNode) OpenDir(
	context *fuse.Context) ([]fuse.DirEntry, fuse.Status) {
	n.log("OpenDir()")
	return []fuse.DirEntry{}, fuse.ENOSYS
}

func (n *loggingNode) GetXAttr(attribute string,
	context *fuse.Context) (data []byte, code fuse.Status) {
	n.log("GetXAttr()")
	return nil, fuse.ENOSYS
}

func (n *loggingNode) RemoveXAttr(attr string,
	context *fuse.Context) fuse.Status {
	n.log("RemoveXAttr()")
	return fuse.ENOSYS
}

func (n *loggingNode) SetXAttr(attr string, data []byte, flags int,
	context *fuse.Context) fuse.Status {
	n.log("SetXAttr()")
	return fuse.ENOSYS
}

func (n *loggingNode) ListXAttr(
	context *fuse.Context) (attrs []string, code fuse.Status) {
	n.log("ListXAttr()")
	return nil, fuse.ENOSYS
}

func (n *loggingNode) GetAttr(out *fuse.Attr, file nodefs.File,
	context *fuse.Context) (code fuse.Status) {
	n.log("GetAttr()")
	return fuse.ENOSYS
}

func (n *loggingNode) Chmod(file nodefs.File, perms uint32,
	context *fuse.Context) (code fuse.Status) {
	n.log("Chmod()")
	return fuse.ENOSYS
}

func (n *loggingNode) Chown(file nodefs.File, uid uint32, gid uint32,
	context *fuse.Context) (code fuse.Status) {
	n.log("Chown()")
	return fuse.ENOSYS
}

func (n *loggingNode) Truncate(file nodefs.File, size uint64,
	context *fuse.Context) (code fuse.Status) {
	n.log("Truncate()")
	return fuse.ENOSYS
}

func (n *loggingNode) Utimens(file nodefs.File, atime *time.Time,
	mtime *time.Time, context *fuse.Context) (code fuse.Status) {
	n.log("Utimens()")
	return fuse.ENOSYS
}

func (n *loggingNode) Fallocate(file nodefs.File, off uint64, size uint64,
	mode uint32, context *fuse.Context) (code fuse.Status) {
	n.log("Fallocate()")
	return fuse.ENOSYS
}

func (n *loggingNode) Read(file nodefs.File, dest []byte, off int64,
	context *fuse.Context) (fuse.ReadResult, fuse.Status) {
	n.log("Read()")
	return nil, fuse.ENOSYS
}

func (n *loggingNode) Write(file nodefs.File, data []byte, off int64,
	context *fuse.Context) (written uint32, code fuse.Status) {
	n.log("Write()")
	return 0, fuse.ENOSYS
}
