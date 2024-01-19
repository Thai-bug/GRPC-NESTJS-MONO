package ENV

import (
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	DBUrl       string
	Environment string
}

var EnvironmentVariables *Environment

func ProcessGetEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	EnvironmentVariables = &Environment{
		DBUrl:       os.Getenv("DB_URL"),
		Environment: os.Getenv("ENV"),
	}
}
