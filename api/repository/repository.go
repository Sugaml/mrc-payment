package repository

import (
	"github.com/Sugaml/mrc-payment/api/models"

	"github.com/jinzhu/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository() *Repository {
	return &Repository{}
}

func Migrate(r *Repository) {
	r.DB.AutoMigrate(
		models.Invoice{},
		models.InvoiceItems{},
		models.Gateway{},
		models.Transaction{},
	)
}
