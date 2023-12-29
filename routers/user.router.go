package routers

import (
	"github.com/gin-gonic/gin"
	"web-service/controllers"
)

func UserRoutes(router *gin.RouterGroup) {
	router.GET("users", controllers.GetUser)
	router.POST("users", controllers.CreateUser)
}
