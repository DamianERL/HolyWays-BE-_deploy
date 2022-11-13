package repositories

import (
	"holyways/models"

	"gorm.io/gorm"
)

type FundRepository interface {
	FindFunds() ([]models.Fund, error)
	FindFundId(ID int) ([]models.Fund, error)
	GetFund(ID int) (models.Fund, error)
	CreateFund(fund models.Fund) (models.Fund, error)
	UpdateFund(fund models.Fund) (models.Fund, error)
	DeleteFund(fund models.Fund) (models.Fund, error)
}

func RepositoryFund(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindFundId(ID int) ([]models.Fund, error) {
	var fundss []models.Fund
	err := r.db.Preload("User").Find(&fundss, "user_id = ?", ID).Error

	return fundss, err
}


func (r *repository) FindFunds() ([]models.Fund, error) {
	var funds []models.Fund
	err := r.db.Preload("User").Find(&funds).Error

	return funds, err
}
func (r *repository) GetFund(ID int) (models.Fund, error) {
	var fund models.Fund
	err := r.db.Preload("User").First(&fund,ID).Error

	return fund, err
}

func (r *repository) CreateFund(fund models.Fund) (models.Fund, error) {
	err := r.db.Create(&fund).Error

	return fund, err
}

func (r *repository) UpdateFund(fund models.Fund) (models.Fund, error) {
	err := r.db.Save(&fund).Error

	return fund, err
}

func (r *repository) DeleteFund(fund models.Fund) (models.Fund, error) {
	err := r.db.Delete(&fund).Error

	return fund, err
}
