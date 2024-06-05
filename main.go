package main

import (
	"shop_banglore/config"
	"shop_banglore/database"

	"github.com/gin-gonic/gin"
)

func init() {
	config.Initialize_env_variables()
	database.Connect()
}

func main() {
	r := gin.Default()

	r.Run()
}
