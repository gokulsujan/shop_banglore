package main

import (
	"shop_banglore/config"

	"github.com/gin-gonic/gin"
)

func init() {
	config.Initialize_env_variables()
	config.ConnectToDb()
}

func main() {
	r := gin.Default()

	r.Run()
}
