package models

import "gorm.io/gorm"

type WalletTransaction struct {
	gorm.Model
	UserId uint `gorm: "not null"`
	User   User `gorm: "ForeignKey:UserId"`
	Amount int  `gorm: "not null"`
}
