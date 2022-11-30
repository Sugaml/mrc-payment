package models

import (
	"github.com/jinzhu/gorm"
)

type Transaction struct {
	gorm.Model
	Title       string   `gorm:"not null" json:"title"`
	User        *User    `gorm:"foreignkey:UserID" json:"user"`
	UserID      uint     `gorm:"not null" json:"user_id"`
	Amount      float64  `gorm:"not null" json:"amount"`
	Invoice     *Invoice `gorm:"foreignkey:InvoiceID" json:"invoice"`
	InvoiceID   uint     `gorm:"not null" json:"invoice_id"`
	Gateway     *Gateway `gorm:"foreignkey:GatewayID" json:"gateway"`
	GatewayID   uint     `gorm:"not null" json:"gateway_id"`
	RefrenceID  string   `gorm:"null" json:"ref_id"`
	PaymentMode string   `gorm:"null" json:"mode"`
	PayLoad     string   `gorm:"null" json:"pay_load"`
	Status      string   `gorm:"not null" json:"status"`
}
