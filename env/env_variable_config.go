package env

import (
	"github.com/joho/godotenv"
)

func LoadFile() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}
