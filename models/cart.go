package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	User      User    `gorm: "ForeignKey:UserId"`
	VarientId uint    `gorm: "not null"`
	Varient   Varient `gorm: "ForeignKey:VarientId"`
	UserId    uint    `gorm: "not null"`
}
