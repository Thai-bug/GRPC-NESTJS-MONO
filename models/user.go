package models

import (
	"time"
	PasswordUtils "web-service/utils/crypto"
	StatusEnum "web-service/utils/enums"

	"gorm.io/gorm"
)

type User struct {
	ID        int64              `json:"id" gorm:"primary_key"`
	Name      string             `json:"name"`
	Email     *string            `json:"email"`
	CreatedAt time.Time          `json:"createdAt"`
	Status    *StatusEnum.Status `json:"status" gorm:"default:1"`
	Password  string             `json:"-"`

	Stores []*Store `json:"stores" -`
	BaseModel
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password, _ = PasswordUtils.HashPassword(u.Password)

	return
}

func (User) TableName() string {
	return "user"
}
