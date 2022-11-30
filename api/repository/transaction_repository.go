package repository

import (
	"errors"

	"github.com/Sugaml/mrc-payment/api/models"

	"github.com/jinzhu/gorm"
)

type ITransaction interface {
	CreateTransaction(db *gorm.DB, data models.Transaction) (models.Transaction, error)
	FilterTransactionByStatus(db *gorm.DB, userID, page, size uint64, status, startdate, enddate, q, sortColumn, sortDirection string) ([]models.Transaction, int64, error)
	FindByIdTransaction(db *gorm.DB, uid uint) (models.Transaction, error)
	FindByInvoiceIdTransaction(db *gorm.DB, pid uint) (models.Transaction, error)
	UpdateTransaction(db *gorm.DB, data models.Transaction) (models.Transaction, error)
	DeleteTransaction(db *gorm.DB, pid uint) (int64, error)
}

type TransactionRepo struct{}

func NewTransactionRepo() ITransaction {
	return &TransactionRepo{}
}

func (d TransactionRepo) CreateTransaction(db *gorm.DB, data models.Transaction) (models.Transaction, error) {
	err := db.Model(&models.Transaction{}).Create(&data).Error
	if err != nil {
		return models.Transaction{}, err
	}
	return data, nil
}

func (d TransactionRepo) FilterTransactionByStatus(db *gorm.DB, userID, page, size uint64, status, startdate, enddate, q, sortColumn, sortDirection string) ([]models.Transaction, int64, error) {
	transactions := []models.Transaction{}
	f := db.Model(&models.Transaction{})
	count := 0
	if len(status) > 0 {
		f = f.Where("status = ? ", status)
	}
	if userID > 0 {
		f = f.Where("user_id= ? ", userID)
	}
	if len(startdate) != 0 && len(enddate) != 0 {
		f = f.Where("created_at BETWEEN ?::timestamp AND ?::timestamp", startdate, enddate)
	}
	f = f.Where("lower(title) LIKE lower(?) OR lower(pay_load) LIKE lower(?)", "%"+q+"%", "%"+q+"%")
	err := f.
		Order(sortColumn + " " + sortDirection).
		Limit(size).
		Offset(size * (page - 1)).
		Preload("User").
		Find(&transactions).Error
	if err != nil {
		return []models.Transaction{}, 0, err
	}
	f.Count(&count)
	return transactions, int64(count), nil
}

func (d TransactionRepo) FindByIdTransaction(db *gorm.DB, uid uint) (models.Transaction, error) {
	data := models.Transaction{}
	err := db.Model(&models.Transaction{}).Where("id = ?", uid).Take(&data).Error
	if err != nil {
		return models.Transaction{}, err
	}
	return data, nil
}

func (d TransactionRepo) FindByInvoiceIdTransaction(db *gorm.DB, pid uint) (models.Transaction, error) {
	data := models.Transaction{}
	err := db.Model(&models.Transaction{}).Where("invoice_id = ?", pid).Take(&data).Error
	if err != nil {
		return models.Transaction{}, err
	}
	return data, nil
}

func (d TransactionRepo) UpdateTransaction(db *gorm.DB, data models.Transaction) (models.Transaction, error) {
	err := db.Model(&models.Transaction{}).Where("id = ?", data.ID).Update(&data).Error
	if err != nil {
		return models.Transaction{}, err
	}
	return data, nil
}

func (d TransactionRepo) DeleteTransaction(db *gorm.DB, pid uint) (int64, error) {
	result := db.Model(&models.Transaction{}).Where("id = ?", pid).Take(&models.Transaction{}).Delete(&models.Transaction{})
	if result.Error != nil {
		if gorm.IsRecordNotFoundError(result.Error) {
			return 0, errors.New("promocode not found")
		}
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
