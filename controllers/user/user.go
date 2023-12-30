package UserController

import (
	"net/http"
	UserDTO "web-service/dto/user"
	"web-service/models"

	StoreService "web-service/services/store"
	UserService "web-service/services/user"
	ResponseError "web-service/utils/errors"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User = UserService.GetUsers()
	var result []models.User

	for _, user := range users {
		mapUser := &user
		mapUser.Stores = []*models.Store{}
		stores := StoreService.GetStoresOfUser(user)
		for _, store := range stores {
			mapUser.Stores = append(user.Stores, &store)
		}
		result = append(result, *mapUser)
	}

	c.JSON(200, result)
}

func CreateUser(c *gin.Context) {
	var input UserDTO.CreateUserDTO

	if err := c.ShouldBindJSON(&input); err != nil {

		customErrors := ResponseError.MapValidationErrors(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": customErrors})

		return
	}

	newUser := UserService.CreateUser(input)

	c.JSON(200, newUser)
}

func Login(c *gin.Context) {

}
