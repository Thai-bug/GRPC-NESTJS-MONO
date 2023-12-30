package routers

import (
	"github.com/gin-gonic/gin"
	AuthController "web-service/controllers/auth"
)

func AuthRoutes(router *gin.RouterGroup) {
	authGroup := router.Group("auth")
	{
		authGroup.POST("login", AuthController.Login)
	}
}
