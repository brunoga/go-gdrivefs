package gdrive

import (
	"fmt"
	"io"
	"net/http"

	"github.com/hanwen/go-fuse/fuse"

	drive "google.golang.org/api/drive/v2"
)

type Handler struct {
	auth         *Auth
	driveService *drive.Service
}

func NewHandler(auth *Auth) (*Handler, error) {
	driveService, err := drive.New(auth.Client())
	if err != nil {
		return nil, err
	}

	return &Handler{
		auth,
		driveService,
	}, nil
}

func (h *Handler) GetFileById(id string) (*drive.File, error) {
	filesService := drive.NewFilesService(h.driveService)

	fileGetCall := filesService.Get(id)

	f, err := fileGetCall.Do()
	if err != nil {
		return nil, err
	}

	return f, nil
}

func (h *Handler) GetFileByName(name, parentId string) (*drive.File, error) {
	filesService := drive.NewFilesService(h.driveService)

	fileListCall := filesService.List()
	fileListCall.Fields("items")
	fileListCall.Q(fmt.Sprintf("'%s' in parents and title = '%s'",
		parentId, name))

	fileList, err := fileListCall.Do()
	if err != nil {
		return nil, err
	}

	if len(fileList.Items) == 0 {
		return nil, nil
	}

	return fileList.Items[0], nil
}

func (h *Handler) GetFileList(parentId string) ([]*drive.File, error) {
	filesService := drive.NewFilesService(h.driveService)

	fileListCall := filesService.List()
	fileListCall.Fields("items")
	fileListCall.Q(fmt.Sprintf("'%s' in parents", parentId))

	fileList, err := fileListCall.Do()
	if err != nil {
		return nil, err
	}

	return fileList.Items, nil
}

func (h *Handler) ReadFile(driveFile *drive.File,
	dest []byte, offset int64) (fuse.ReadResult, fuse.Status) {
	if len(driveFile.DownloadUrl) == 0 {
		return nil, fuse.ENODATA
	}

	c := h.auth.Client()

	req, err := http.NewRequest("GET", driveFile.DownloadUrl, nil)
	if err != nil {
		return nil, fuse.EIO
	}

	req.Header.Add("Range", fmt.Sprintf("bytes=%d-%d", offset,
		offset+int64(len(dest))))

	resp, err := c.Do(req)
	if err != nil {
		return nil, fuse.EIO
	}
	defer resp.Body.Close()

	tRead := 0
	for {
		nRead, err := resp.Body.Read(dest[tRead:])

		tRead += nRead

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, fuse.EIO
		}

		if tRead == len(dest) {
			break
		}
	}

	return fuse.ReadResultData(dest[:tRead]), fuse.OK
}
