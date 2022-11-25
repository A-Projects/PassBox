package folder

import (
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"passbox.server/domain"
)

type service struct {
	logger *logrus.Logger
}

func (s *service) GetFolders() ([]domain.FolderEntity, error) {
	s.logger.Info("folders list")

	folders := []domain.FolderEntity{}

	id, _ := uuid.NewV4()
	folders = append(folders, domain.FolderEntity{
		Id:   id,
		Name: "Папка 1",
	})

	id, _ = uuid.NewV4()
	folders = append(folders, domain.FolderEntity{
		Id:   id,
		Name: "Папка 2",
	})

	id, _ = uuid.NewV4()
	folders = append(folders, domain.FolderEntity{
		Id:   id,
		Name: "Папка 3",
	})

	return folders, nil
}
