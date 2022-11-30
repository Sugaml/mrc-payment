package repository

import (
	"errors"

	"github.com/Sugaml/mrc-payment/api/models"

	"github.com/jinzhu/gorm"
)

type IGateway interface {
	CreateGateway(db *gorm.DB, data models.Gateway) (models.Gateway, error)
	FindAllGateway(db *gorm.DB) ([]models.Gateway, error)
	FindActiveGateway(db *gorm.DB) ([]models.Gateway, error)
	FindByIdGateway(db *gorm.DB, id uint) (models.Gateway, error)
	FindByCodeGateway(db *gorm.DB, code string) (models.Gateway, error)
	UpdateGateway(db *gorm.DB, data models.Gateway) (models.Gateway, error)
	DeleteGateway(db *gorm.DB, pid uint) (int64, error)
}

type GatewayRepo struct{}

func NewGetwayRepo() IGateway {
	return &GatewayRepo{}
}

func (d GatewayRepo) CreateGateway(db *gorm.DB, data models.Gateway) (models.Gateway, error) {
	err := db.Model(&models.Gateway{}).Create(&data).Error
	if err != nil {
		return models.Gateway{}, err
	}
	return data, nil
}

func (d GatewayRepo) FindAllGateway(db *gorm.DB) ([]models.Gateway, error) {
	datas := []models.Gateway{}
	err := db.Model(&models.Gateway{}).Order("id desc").Find(&datas).Error
	if err != nil {
		return []models.Gateway{}, err
	}
	return datas, nil
}

func (d GatewayRepo) FindActiveGateway(db *gorm.DB) ([]models.Gateway, error) {
	datas := []models.Gateway{}
	err := db.Model(&models.Gateway{}).Where("active = ?", true).Order("id desc").Find(&datas).Error
	if err != nil {
		return []models.Gateway{}, err
	}
	return datas, nil
}

func (d GatewayRepo) FindByIdGateway(db *gorm.DB, id uint) (models.Gateway, error) {
	data := models.Gateway{}
	err := db.Model(&models.Gateway{}).Where("id = ?", id).Take(&data).Error
	if err != nil {
		return models.Gateway{}, err
	}
	return data, nil
}

func (d GatewayRepo) FindByCodeGateway(db *gorm.DB, code string) (models.Gateway, error) {
	data := models.Gateway{}
	err := db.Model(&models.Gateway{}).Where("code = ?", code).Take(&data).Error
	if err != nil {
		return models.Gateway{}, err
	}
	return data, nil
}

func (d GatewayRepo) UpdateGateway(db *gorm.DB, data models.Gateway) (models.Gateway, error) {
	gateway := map[string]interface{}{
		"active": data.Active,
	}
	if data.InitUrl != "" {
		gateway["init_url"] = data.InitUrl
	}
	if data.Name != "" {
		gateway["name"] = data.Name
	}
	if data.SuccessUrl != "" {
		gateway["success_url"] = data.SuccessUrl
	}
	if data.FailureUrl != "" {
		gateway["failure_url"] = data.FailureUrl
	}
	if data.PaymentUrl != "" {
		gateway["payment_url"] = data.PaymentUrl
	}
	if data.Description != "" {
		gateway["description"] = data.Description
	}
	if data.Code != "" {
		gateway["code"] = data.Code
	}
	if data.Country != "" {
		gateway["country"] = data.Country
	}
	if data.Key.RawMessage != nil {
		gateway["key"] = data.Key
	}
	err := db.Model(&models.Gateway{}).Where("id = ?", data.ID).Updates(gateway).Error
	if err != nil {
		return models.Gateway{}, err
	}
	return data, nil
}

func (d GatewayRepo) DeleteGateway(db *gorm.DB, pid uint) (int64, error) {
	result := db.Model(&models.Gateway{}).Where("id = ?", pid).Take(&models.Gateway{}).Delete(&models.Gateway{})
	if result.Error != nil {
		if gorm.IsRecordNotFoundError(result.Error) {
			return 0, errors.New("gateway not found")
		}
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
