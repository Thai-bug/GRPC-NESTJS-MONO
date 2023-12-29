package connections

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"web-service/models"
)

var DB *gorm.DB
var ConnectionError error

func ConnectDB() {
	db, error := gorm.Open(postgres.Open("postgres://thai-bug:12022021@localhost:9013?database=golang"), &gorm.Config{})

	db.AutoMigrate(&models.User{})

	DB = db
	ConnectionError = error
}
