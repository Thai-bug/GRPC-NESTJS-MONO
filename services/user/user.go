package UserService

import (
	"web-service/connections"
	UserDTO "web-service/dto/user"
	"web-service/models"
)

func GetUsers() []models.User {
	var result []models.User
	connections.DB.Find(&result)

	return result
}

func CreateUser(data UserDTO.CreateUserDTO) models.User {
	user := models.User{Name: data.Name, Email: &data.Email, Password: data.Password}
	connections.DB.Create(&user)

	return user
}

func FindUserById(id int64) models.User {
	var user models.User
	connections.DB.First(&user, id)

	return user
}
