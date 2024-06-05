package main

import (
	env "shop_banglore/config"
	"shop_banglore/database"

	"github.com/gin-gonic/gin"
)

func init() {
	env.LoadFile()
	database.Connect()
}

func main() {
	r := gin.Default()

	r.Run()
}
