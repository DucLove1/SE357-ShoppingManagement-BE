package route

import (
	"github.com/DucLove1/SE357-ShoppingManagement-BE/internal/controller"
	"github.com/DucLove1/SE357-ShoppingManagement-BE/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterMediaRoutes(rg *gin.RouterGroup, c *controller.MediaController) {
	media := rg.Group("/media")
	media.Use(middleware.RequireAuth())
	{
		// Upload
		media.POST("/image", c.UploadImage)
		media.POST("/images", c.UploadImages)
		media.POST("/video", c.UploadVideo)
		media.POST("/videos", c.UploadVideos)

	}
}
