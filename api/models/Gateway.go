package models

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

type Gateway struct {
	gorm.Model
	Name        string         `gorm:"not null" json:"name"`
	Code        string         `gorm:"not null" json:"code" `
	Icon        string         `gorm:"null" json:"icon" `
	Description string         `gorm:"null" json:"description" `
	Country     string         `gorm:"null" json:"country" `
	InitUrl     string         `gorm:"not null" json:"init_url"`
	PaymentUrl  string         `gorm:"not null" json:"payment_url"`
	SuccessUrl  string         `gorm:"not null" json:"success_url"`
	FailureUrl  string         `gorm:"not null" json:"failure_url"`
	Active      bool           `gorm:"default:true" json:"active"`
	Key         postgres.Jsonb `gorm:"not null" json:"key"`
}
