package gdrive

import (
	"fmt"

	drive "google.golang.org/api/drive/v2"
)

type Handler struct {
	auth         *Auth
	driveService *drive.Service
}

type handlerResult struct {
	gDriveFiles []*drive.File
	gDriveError error
}

func (r *handlerResult) Get() ([]*drive.File, error) {
	return r.gDriveFiles, r.gDriveError
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

func (h *Handler) GetFileById(id string) <-chan *handlerResult {
	c := make(chan *handlerResult)

	go h.internalGetFileById(id, c)

	return c
}

func (h *Handler) internalGetFileById(id string, c chan<- *handlerResult) {
	defer close(c)

	filesService := drive.NewFilesService(h.driveService)
	fileGetCall := filesService.Get(id)
	f, err := fileGetCall.Do()
	if err != nil {
		c <- &handlerResult{
			nil,
			err,
		}

		return
	}

	c <- &handlerResult{
		[]*drive.File{
			f,
		},
		nil,
	}
}

func (h *Handler) GetFileByName(name, parentId string) <-chan *handlerResult {
	c := make(chan *handlerResult)

	go h.internalGetFileByName(name, parentId, c)

	return c
}

func (h *Handler) internalGetFileByName(name, parentId string, c chan<- *handlerResult) {
	defer close(c)

	filesService := drive.NewFilesService(h.driveService)
	fileListCall := filesService.List()
	fileListCall.Fields("items")
	fileListCall.Q(fmt.Sprintf("'%s' in parents and title = '%s'",
		parentId, name))
	fileList, err := fileListCall.Do()
	if err != nil {
		c <- &handlerResult{
			nil,
			err,
		}

		return
	}

	c <- &handlerResult{
		fileList.Items,
		nil,
	}
}

func (h *Handler) GetFileList(parentId string) <-chan *handlerResult {
	c := make(chan *handlerResult)

	go h.internalGetFileList(parentId, c)

	return c
}

func (h *Handler) internalGetFileList(parentId string, c chan<- *handlerResult) {
	defer close(c)

	filesService := drive.NewFilesService(h.driveService)
	fileListCall := filesService.List()
	fileListCall.Fields("items")
	fileListCall.Q(fmt.Sprintf("'%s' in parents", parentId))
	fileList, err := fileListCall.Do()
	if err != nil {
		c <- &handlerResult{
			nil,
			err,
		}

		return
	}

	c <- &handlerResult{
		fileList.Items,
		nil,
	}
}
