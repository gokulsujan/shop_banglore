package config

import (
	"github.com/joho/godotenv"
)

func Initialize_env_variables() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}
