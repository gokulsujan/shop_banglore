package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserId    uint    `gorm: "not null"`
	User      User    `gorm: "ForeignKey: UserId"`
	PaymentId uint    `gorm: "not null"`
	Payment   Payment `gorm: "ForeignKey: PaymentId"`
}
