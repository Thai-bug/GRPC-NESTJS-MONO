package routers

import (
	"github.com/gin-gonic/gin"
	UserController "web-service/controllers/user"
)

func UserRoutes(router *gin.RouterGroup) {
	group := router.Group("users")
	{
		group.GET("", UserController.GetUsers)
		group.POST("", UserController.CreateUser)
	}
}
