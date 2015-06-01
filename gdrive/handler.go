package gdrive

import (
	"fmt"

	drive "google.golang.org/api/drive/v2"
)

type Handler struct {
	auth         *Auth
	driveService *drive.Service
}

type HandlerResult struct {
	gDriveFiles []*drive.File
	gDriveError error
}

func (r *HandlerResult) GetDriveFiles() []*drive.File {
	return r.gDriveFiles
}

func (r *HandlerResult) GetDriveError() error {
	return r.gDriveError
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

func (h *Handler) GetFileById(id string) <-chan *HandlerResult {
	c := make(chan *HandlerResult)

	go h.internalGetFileById(id, c)

	return c
}

func (h *Handler) internalGetFileById(id string, c chan<- *HandlerResult) {
	defer close(c)

	filesService := drive.NewFilesService(h.driveService)
	fileGetCall := filesService.Get(id)
	f, err := fileGetCall.Do()
	if err != nil {
		c <- &HandlerResult{
			nil,
			err,
		}

		return
	}

	c <- &HandlerResult{
		[]*drive.File{
			f,
		},
		nil,
	}
}

func (h *Handler) GetFileByName(name, parentId string) <-chan *HandlerResult {
	c := make(chan *HandlerResult)

	go h.internalGetFileByName(name, parentId, c)

	return c
}

func (h *Handler) internalGetFileByName(name, parentId string, c chan<- *HandlerResult) {
	defer close(c)

	filesService := drive.NewFilesService(h.driveService)
	fileListCall := filesService.List()
	fileListCall.Fields("items")
	fileListCall.Q(fmt.Sprintf("'%s' in parents and title = '%s'",
		parentId, name))
	fileList, err := fileListCall.Do()
	if err != nil {
		c <- &HandlerResult{
			nil,
			err,
		}

		return
	}

	c <- &HandlerResult{
		fileList.Items,
		nil,
	}
}

func (h *Handler) GetFileList(parentId string) <-chan *HandlerResult {
	c := make(chan *HandlerResult)

	go h.internalGetFileList(parentId, c)

	return c
}

func (h *Handler) internalGetFileList(parentId string, c chan<- *HandlerResult) {
	defer close(c)

	filesService := drive.NewFilesService(h.driveService)
	fileListCall := filesService.List()
	fileListCall.Fields("items")
	fileListCall.Q(fmt.Sprintf("'%s' in parents", parentId))
	fileList, err := fileListCall.Do()
	if err != nil {
		c <- &HandlerResult{
			nil,
			err,
		}

		return
	}

	c <- &HandlerResult{
		fileList.Items,
		nil,
	}
}
