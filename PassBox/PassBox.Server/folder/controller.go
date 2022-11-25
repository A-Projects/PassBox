package folder

import (
	"github.com/gin-gonic/gin"
	"passbox.server/domain"
)

type controller struct {
	folderService domain.FolderService
}

// GetFolders godoc
//
//	@Summary		Список папок
//	@Description	Предоставляет список папок
//	@Tags			folders
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		domain.Folder
//	@Failure		400	{object}	helper.HTTPError
//	@Failure		404	{object}	helper.HTTPError
//	@Failure		500	{object}	helper.HTTPError
//	@Router			/folders/ [get]
func (c *controller) GetFolders(ctx *gin.Context) {
	folders, _ := c.folderService.GetFolders()
	ctx.JSON(200, folders)
}
