package UserDTO

type GetUserQuery struct {
	ID string `form:"id" json:"id"`
}

type CreateUserDTO struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
