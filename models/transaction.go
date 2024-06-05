package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	PaymentId uint    `gorm: "not null"`
	Payment   Payment `gorm: "ForeignKey:PaymentId"`
	Amount    uint    `gorm: "not null"`
	Method    uint    `gorm: "not null"`  //1=>UPI 2=>Card 3=>Netbanking 4=>Wallet
	Status    uint    `gorm: "default:1"` // 1=>Pending 2=>Success 3=>Reverted
}
