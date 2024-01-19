package connections

import (
	"log"
	"os"
	"time"
	"web-service/models"

	ENV "web-service/utils/env"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var ConnectionError error

func ConnectDB() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  false,         // Disable color
		},
	)

	db, error := gorm.Open(postgres.Open(ENV.EnvironmentVariables.DBUrl), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   newLogger,
	})

	if ENV.EnvironmentVariables.Environment != "production" {
		db.AutoMigrate(&models.User{})
		db.AutoMigrate(&models.Store{})
	}

	DB = db
	ConnectionError = error
}
