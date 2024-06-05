package env

import (
	"github.com/joho/godotenv"
)

func InitializeVariables() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}
