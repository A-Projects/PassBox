package domain

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/gofrs/uuid"
)

// Папка
// @Description Папка
type Folder struct {
	// Идентификатор
	Id uuid.UUID `json:"id" example:"9b4dc8d9-6da7-40c5-b7b1-30e81edca64d"`

	// Название
	Name string `json:"name" example:"folder name"`
}

// Папка (для создания)
// @Description Папка (для создания)
type FolderForCreate struct {
	// Название
	Name string `json:"name" example:"folder name"`
}

// Папка
type FolderEntity struct {
	// Идентификатор
	Id uuid.UUID

	// Название
	Name string
}

type FolderService interface {
	GetFolders() ([]FolderEntity, error)
}

type FolderController interface {
	GetFolders(ctx *gin.Context)
}
