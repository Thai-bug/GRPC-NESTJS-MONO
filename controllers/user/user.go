package UserController

import (
	"net/http"
	UserDTO "web-service/dto/user"
	"web-service/models"

	"github.com/gin-gonic/gin"
	UserService "web-service/services/user"
	ResponseError "web-service/utils/errors"
)

func GetUsers(c *gin.Context) {
	var users []models.User = UserService.GetUsers()

	c.JSON(200, users)
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
