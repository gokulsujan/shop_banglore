package main

import (
	"shop_banglore/database"
	"shop_banglore/env"

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
