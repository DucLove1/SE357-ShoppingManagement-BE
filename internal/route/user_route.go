package route

import (
	"github.com/DucLove1/SE357-ShoppingManagement-BE/internal/controller"
	"github.com/DucLove1/SE357-ShoppingManagement-BE/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup, c *controller.UserController) {
	users := rg.Group("/users")

	users.Use(middleware.RequireAuth(), middleware.RequireAdmin())
	{
		users.GET("buyers", c.GetBuyers)
		users.GET("buyers/", c.GetBuyers)
		users.GET("sellers", c.GetSellers)
		users.GET("sellers/", c.GetSellers)
	}
	
	// Routes for the currently authenticated user ("me")
	me := users.Group("/me")
	me.Use(middleware.RequireAuth())
	{
		me.GET("", c.GetMyProfile)
		me.PUT("/profile", c.UpdateProfile)
		me.PUT("/password", c.ChangePassword)
	}
}
