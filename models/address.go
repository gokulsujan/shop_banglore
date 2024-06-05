package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	Line1   string `gorm: "not null"`
	Line2   string
	Line3   string
	State   string `gorm: "not null"`
	Pincode string `gorm: "not null"`
}
