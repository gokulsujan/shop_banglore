package config

import (
	"os"
	"shop_banglore/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	dsn := os.Getenv("dsn")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	DB.AutoMigrate(
									&models.User{},
									&models.Address{},
									&models.WalletTransaction{},
									&models.Category{},
									&models.Product{},
									&models.Varient{},
									&models.Cart{},
									&models.Payment{},
									&models.Order{},
									&models.OrderItem{},
									&models.Transaction{},
								)
}
