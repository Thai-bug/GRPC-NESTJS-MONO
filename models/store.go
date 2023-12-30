package models

import (
	StatusEnum "web-service/utils/enums"
	StringUtils "web-service/utils/string"

	"gorm.io/gorm"
)

type Store struct {
	ID     int16             `json:"id"`
	Name   string            `json:"name"`
	Status StatusEnum.Status `json:"status" gorm:"default:1"`
	UserId int64             `json:"userId"`
	Slug   string            `json:"slug"`

	User *User `json:"user" -`
}

func (u *Store) BeforeCreate(tx *gorm.DB) (err error) {
	u.Slug = StringUtils.GenerateSlug(u.Name)

	return
}

func (s *Store) TableName() string {
	return "store"
}
