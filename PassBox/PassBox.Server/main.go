package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "passbox.server/docs"
	"passbox.server/folder"
)

// @title          PassBox API 1.0.
// @version        1.0
// @description    Данное API реализует функции управления паролями.

// @license.name GNU GPL v3.0
// @license.url  https://www.gnu.org/licenses/gpl-3.0.html

// @host     localhost:5001
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})

	router := gin.Default()

	folderController := folder.Wire(logger)

	v1 := router.Group("/api/v1")
	{

		folders := v1.Group("/folders")
		{
			folders.GET("", folderController.GetFolders)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":5001")
}

/*
func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.GetHeader("Authorization")) == 0 {
			helper.NewError(c, http.StatusUnauthorized, errors.New("authorization is required Header"))
			c.Abort()
		}
		c.Next()
	}
}*/
