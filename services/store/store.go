package StoreService

import (
	"errors"
	"web-service/connections"
	StoreDTO "web-service/dto/store"
	"web-service/models"
	StringUtils "web-service/utils/string"
)

func CreateStore(data StoreDTO.CreateStoreDTO) (models.Store, error) {
	existedSlug := StringUtils.GenerateSlug(data.Name)
	existedStore := GetStoreBySlug(existedSlug)
	var store models.Store

	if existedStore != (models.Store{}) {
		return store, errors.New("store already exists")
	}

	store = models.Store{Name: data.Name, UserId: data.UserId}

	connections.DB.Create(&store)

	return store, nil
}

func GetStoreBySlug(slug string) models.Store {
	var store models.Store
	connections.DB.Where("slug = ? and status = 1", slug).First(&store)

	return store
}

func GetStoresOfUser(user models.User) []models.Store {
	var stores []models.Store
	connections.DB.Where("user_id = ? and status = 1", user.ID).Find(&stores)

	return stores
}

func GetStores() []models.Store {
	var stores []models.Store
	connections.DB.Where("status = 1").Find(&stores)

	return stores
}
