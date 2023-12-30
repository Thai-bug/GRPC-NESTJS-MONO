package models

import "time"

type BaseModel struct {
	CreatedAt time.Time `gorm:"->:false;column:created_at" json:"-"`
	UpdatedAt time.Time `gorm:"->:false;column:updated_at" json:"-"`
}

type Tabler interface {
	TableName() string
}
