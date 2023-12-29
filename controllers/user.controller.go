package controllers

import (
	"fmt"
	"net/http"
	UserDTO "web-service/dto"
	"web-service/models"
	UserService "web-service/services/user"
	ResponseError "web-service/utils/errors"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	var users []models.User = UserService.GetUsers()

	fmt.Println("User:", users)

	c.JSON(200, users)
}

func CreateUser(c *gin.Context) {
	var input *UserDTO.CreateUserDTO

	if err := c.ShouldBindJSON(&input); err != nil {

		customErrors := ResponseError.MapValidationErrors(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": customErrors})
		return
	}

	UserService.CreateUser(input)

	c.JSON(200, input)
}
