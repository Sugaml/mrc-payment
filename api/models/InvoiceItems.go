package models

import (
	"github.com/jinzhu/gorm"
)

type InvoiceItems struct {
	gorm.Model
	Invoice    *Invoice `gorm:"foreignkey:InvoiceID" json:"invoice"`
	InvoiceID  uint     `gorm:"not null" json:"invoice_id"`
	User       *User    `gorm:"foreignkey:UserID" json:"user"`
	UserID     uint     `gorm:"not null" json:"user_id"`
	Particular string   `gorm:"not null" json:"particular"`
	Rate       float64  `gorm:"not null" json:"rate"`
	Quantity   uint     `gorm:"not null" json:"quantity"`
	Total      float64  `gorm:"not null" json:"total"`
}
