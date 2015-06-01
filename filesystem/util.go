package filesystem

import (
	"fmt"
	"time"

	"github.com/hanwen/go-fuse/fuse"

	drive "google.golang.org/api/drive/v2"
)

func fillAttr(n *loggingNode, driveFile *drive.File, attr *fuse.Attr) bool {
	isDir := false

	// Set file size.
	attr.Size = uint64(driveFile.FileSize)

	// Set created and modified dates.
	createdDate, err := time.Parse(time.RFC3339, driveFile.CreatedDate)
	if err == nil {
		attr.Ctime = uint64(createdDate.Unix())
	} else {
		n.log(fmt.Sprintf("Error parsing date : %s", err))
	}

	modifiedDate, err := time.Parse(time.RFC3339, driveFile.ModifiedDate)
	if err == nil {
		attr.Mtime = uint64(modifiedDate.Unix())
	} else {
		n.log(fmt.Sprintf("Error parsing date : %s", err))
	}

	// Set up file/directory entry.
	if driveFile.MimeType == "application/vnd.google-apps.folder" {
		// This is a directory.
		attr.Mode = fuse.S_IFDIR | 0755
		isDir = true
	} else {
		// This is a file.
		attr.Mode = fuse.S_IFREG | 0644
		isDir = false
	}

	return isDir
}
