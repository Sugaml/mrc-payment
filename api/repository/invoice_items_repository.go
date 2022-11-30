package repository

import (
	"errors"

	"github.com/Sugaml/mrc-payment/api/models"

	"github.com/jinzhu/gorm"
)

type IInvoiceItems interface {
	CreateInvoiceItems(db *gorm.DB, data models.InvoiceItems) (models.InvoiceItems, error)
	FindAllInvoiceItems(db *gorm.DB) ([]models.InvoiceItems, error)
	FindByIdInvoiceItems(db *gorm.DB, uid uint) (models.InvoiceItems, error)
	UpdateInvoiceItems(db *gorm.DB, data models.InvoiceItems) (models.InvoiceItems, error)
	DeleteInvoiceItems(db *gorm.DB, pid uint) (int64, error)
	FindByInvoiceIdInvoiceItems(db *gorm.DB, pid, uid uint) ([]models.InvoiceItems, error)
}

type InvoiceItemRepo struct{}

func NewInvoiceItemRepo() IInvoiceItems {
	return &InvoiceItemRepo{}
}

func (d InvoiceItemRepo) CreateInvoiceItems(db *gorm.DB, data models.InvoiceItems) (models.InvoiceItems, error) {
	err := db.Model(&models.InvoiceItems{}).Create(&data).Error
	if err != nil {
		return models.InvoiceItems{}, err
	}
	return data, nil
}

func (d InvoiceItemRepo) FindAllInvoiceItems(db *gorm.DB) ([]models.InvoiceItems, error) {
	datas := []models.InvoiceItems{}
	err := db.Model(&models.InvoiceItems{}).Order("id desc").Find(&datas).Error
	if err != nil {
		return []models.InvoiceItems{}, err
	}
	return datas, nil
}

func (d InvoiceItemRepo) FindByIdInvoiceItems(db *gorm.DB, uid uint) (models.InvoiceItems, error) {
	data := models.InvoiceItems{}
	err := db.Model(&models.InvoiceItems{}).Preload("Invoice").Where("id = ?", uid).Take(&data).Error
	if err != nil {
		return models.InvoiceItems{}, err
	}
	return data, nil
}

func (d InvoiceItemRepo) FindByInvoiceIdInvoiceItems(db *gorm.DB, pid, uid uint) ([]models.InvoiceItems, error) {
	data := []models.InvoiceItems{}
	err := db.Model(&models.InvoiceItems{}).Where("invoice_id = ? and user_id=?", pid, uid).Find(&data).Error
	if err != nil {
		return []models.InvoiceItems{}, err
	}
	return data, nil
}

func (d InvoiceItemRepo) UpdateInvoiceItems(db *gorm.DB, data models.InvoiceItems) (models.InvoiceItems, error) {
	var invoiceItems = models.InvoiceItems{}
	if data.InvoiceID != 0 {
		invoiceItems.InvoiceID = data.InvoiceID
	}
	if data.UserID != 0 {
		invoiceItems.UserID = data.UserID
	}
	if data.Particular != "" {
		invoiceItems.Particular = data.Particular
	}
	if data.Rate != 0 {
		invoiceItems.Rate = data.Rate
	}
	if data.Total != 0 {
		invoiceItems.Total = data.Total
	}
	err := db.Model(&models.InvoiceItems{}).Where("id=?", data.ID).Updates(invoiceItems).Error
	if err != nil {
		return models.InvoiceItems{}, err
	}
	return data, nil
}

func (d InvoiceItemRepo) DeleteInvoiceItems(db *gorm.DB, pid uint) (int64, error) {
	result := db.Model(&models.InvoiceItems{}).Where("id = ?", pid).Take(&models.InvoiceItems{}).Delete(&models.InvoiceItems{})
	if result.Error != nil {
		if gorm.IsRecordNotFoundError(result.Error) {
			return 0, errors.New("InvoiceItems not found")
		}
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
