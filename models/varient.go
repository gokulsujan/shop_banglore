package models

import "gorm.io/gorm"

type Varient struct {
	gorm.Model
	ProductId uint `gorm: "not null"`
	Product Product `gorm: "ForeignKey: ProductId"`
	Name string
	Status uint `gorm: "not null"`
}
