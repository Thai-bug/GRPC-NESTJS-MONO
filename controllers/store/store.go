package StoreController

import (
	"net/http"
	"web-service/dto/store"
	"web-service/models"
	StoreService "web-service/services/store"
	UserService "web-service/services/user"
	ResponseError "web-service/utils/errors"

	"github.com/gin-gonic/gin"
)

func CreateStore(c *gin.Context) {
	var input StoreDTO.CreateStoreDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		customErrors := ResponseError.MapValidationErrors(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": customErrors})

		return
	}

	existedUser := UserService.FindUserById(input.UserId)
	if existedUser.ID == 0 {
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}

	newStore, error := StoreService.CreateStore(input)
	if error != nil {
		c.JSON(400, gin.H{"error": error.Error()})
		return
	}

	newStore.User = &existedUser
	c.JSON(200, newStore)
}

func GetStores(c *gin.Context) {
	var stores []models.Store = StoreService.GetStores()

	c.JSON(200, stores)
}
