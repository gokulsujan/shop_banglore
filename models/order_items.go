package models

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	OrderId   uint    `gorm: "not null"`
	Order     Order   `gorm: "ForeignKey:OrderId"`
	VarientId uint    `gorm: "not null"`
	Varient   Varient `gorm: "ForeignKey:VarientId"`
}
