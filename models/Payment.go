package models

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	Amount int `gorm: "not null"`
	Type   int `gorm: "not null"` // 1=>COD, 2=>Online
	Status int `gorm: "default:1"` // 1=>Pending 2=>Success 4=>Refunded
	
}
