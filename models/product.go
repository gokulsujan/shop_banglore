package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	CategoryId  uint     `gorm: "not null`
	Category    Category `gorm: "ForeignKey:CategoryId"`
	Description string
	Status      uint `gorm: "not null"`
}
