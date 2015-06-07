package gdrive

import (
	"fmt"

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
