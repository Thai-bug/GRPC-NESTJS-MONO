package StoreDTO

type CreateStoreDTO struct {
	Name   string `json:"name" binding:"required"`
	UserId int64  `json:"userId" binding:"required"`
}
