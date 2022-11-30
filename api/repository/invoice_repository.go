package repository

import (
	"errors"

	"github.com/Sugaml/mrc-payment/api/models"

	"github.com/jinzhu/gorm"
)

type IInvoice interface {
	CreateInvoice(db *gorm.DB, data models.Invoice) (models.Invoice, error)
	FindByIdInvoice(db *gorm.DB, uid uint) (models.Invoice, error)
	FilterInvoiceByStatus(db *gorm.DB, userID, page, size uint64, status, startdate, enddate, q, sortColumn, sortDirection string) ([]models.Invoice, int64, error)
	FindByUserIdInvoice(db *gorm.DB, pid uint) (models.Invoice, error)
	FindByUserIdLastPaidInvoice(db *gorm.DB, pid uint) (models.Invoice, error)
	UpdateInvoice(db *gorm.DB, data models.Invoice) (models.Invoice, error)
	DeleteInvoice(db *gorm.DB, pid uint) (int64, error)
}

type InvoiceRepo struct{}

func NewInvoiceRepo() IInvoice {
	return &InvoiceRepo{}
}

func (d *InvoiceRepo) CreateInvoice(db *gorm.DB, data models.Invoice) (models.Invoice, error) {
	err := db.Model(&models.Invoice{}).Create(&data).Error
	if err != nil {
		return models.Invoice{}, err
	}
	return data, nil
}

func (d *InvoiceRepo) FilterInvoiceByStatus(db *gorm.DB, userID, page, size uint64, status, startdate, enddate, q, sortColumn, sortDirection string) ([]models.Invoice, int64, error) {
	invoices := []models.Invoice{}
	f := db.Model(&models.Invoice{}).Preload("InvoiceItems").Preload("Deduction").Preload("PromoCode")
	count := 0
	if status == "true" {
		f = f.Where("is_paid = ? ", true)
	}
	if status == "false" {
		f = f.Where("is_paid = ? ", false)
	}
	if userID > 0 {
		f = f.Where("user_id= ? ", userID)
	}
	if len(startdate) != 0 && len(enddate) != 0 {
		f = f.Where("created_at BETWEEN ?::timestamp AND ?::timestamp", startdate, enddate)
	}
	f = f.Where("lower(title) LIKE lower(?) OR lower(remarks) LIKE lower(?)", "%"+q+"%", "%"+q+"%")
	err := f.
		Order(sortColumn + " " + sortDirection).
		Limit(size).
		Offset(size * (page - 1)).
		Preload("User").
		Find(&invoices).Error
	if err != nil {
		return []models.Invoice{}, 0, err
	}
	f.Count(&count)
	return invoices, int64(count), nil
}

func (d *InvoiceRepo) FindByIdInvoice(db *gorm.DB, pid uint) (models.Invoice, error) {
	data := models.Invoice{}
	err := db.Model(&models.Invoice{}).Preload("InvoiceItems").Preload("Deduction").Preload("PromoCode").Where("id = ?", pid).Find(&data).Error
	if err != nil {
		return models.Invoice{}, err
	}
	return data, nil
}

func (d *InvoiceRepo) FindByUserIdInvoice(db *gorm.DB, pid uint) (models.Invoice, error) {
	data := models.Invoice{}
	err := db.Model(&models.Invoice{}).Where("user_id = ?", pid).Order("id desc").Take(&data).Error
	if err != nil {
		return models.Invoice{}, err
	}
	return data, nil
}

func (d *InvoiceRepo) FindByUserIdLastPaidInvoice(db *gorm.DB, pid uint) (models.Invoice, error) {
	data := models.Invoice{}
	err := db.Model(&models.Invoice{}).Where("user_id = ? and is_paid = true", pid).Order("id desc").Take(&data).Error
	if err != nil {
		return models.Invoice{}, err
	}
	return data, nil
}
func (d *InvoiceRepo) UpdateInvoice(db *gorm.DB, data models.Invoice) (models.Invoice, error) {
	err := db.Model(&models.Invoice{}).Where("id = ?", data.ID).Update(&data).Error
	if err != nil {
		return models.Invoice{}, err
	}
	return data, nil
}

func (d *InvoiceRepo) DeleteInvoice(db *gorm.DB, id uint) (int64, error) {
	result := db.Model(&models.Invoice{}).Where("id = ?", id).Take(&models.Invoice{}).Delete(&models.Invoice{})
	if result.Error != nil {
		if gorm.IsRecordNotFoundError(result.Error) {
			return 0, errors.New("invoice not found")
		}
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
