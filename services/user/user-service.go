package UserService

import (
	"fmt"
	"web-service/connections"
	UserDTO "web-service/dto"
	"web-service/models"
)

func GetUsers() []models.User {
	var result []models.User
	connections.DB.Find(&result)

	for _, user := range result {
		fmt.Println(user.ID, user.Name)
	}

	return result
}

func CreateUser(user *UserDTO.CreateUserDTO) {
	connections.DB.Create(&user)
}
