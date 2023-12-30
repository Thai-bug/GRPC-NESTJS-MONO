package routers

import (
	"github.com/gin-gonic/gin"
	StoreController "web-service/controllers/store"
)

func StoreRoutes(router *gin.RouterGroup) {
	group := router.Group("stores")
	{
		group.POST("", StoreController.CreateStore)
		group.GET("", StoreController.GetStores)
	}
}
