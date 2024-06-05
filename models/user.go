package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name         string `gorm: "not null"`
	MobileNumber string `gorm: "not null"`
	Email        string `gorm: "not null"`
	Password     string `gorm: "not null"`
	Status       uint   `gorm: "default:1"` // 1=> Pending, 2=>Active, 3=> Inactive, 4=> Disabled
}
