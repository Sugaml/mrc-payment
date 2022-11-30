package models

import (
	"github.com/jinzhu/gorm"
)

type Invoice struct {
	gorm.Model
	Title        string         `gorm:"null" json:"title"`
	Remarks      string         `gorm:"null" json:"remarks"`
	User         *User          `gorm:"foreignkey:UserID" json:"user"`
	UserID       uint           `gorm:"not null" json:"user_id"`
	SubTotal     float64        `gorm:"default:0" json:"sub_total"`
	TotalCost    float64        `gorm:"default:0" json:"total_cost"`
	IsPaid       bool           `gorm:"default:false" json:"is_paid"`
	InvoiceItems []InvoiceItems `gorm:"ForeignKey:InvoiceID" json:"invoice_items"`
}
