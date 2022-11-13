package repositories

import (
	"holyways/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindTransactions(ID int) ([]models.Transaction, error)
	FindTransactionall(FundID int) (models.Transaction, error)
	FindTransactionss(FundID int) ([]models.Transaction, error)
	FindTransactionx(FundID int, UserFundID int) ([]models.Transaction, error)
	GetTransaction(ID int) (models.Transaction, error)
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
	GetOneTransaction(ID string) (models.Transaction, error)
	UpdateTransaction(status string, ID string,Fund models.Fund,donate int) error	
	GetFundTransaction(ID int) (models.Fund, error)
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTransactions(ID int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Preload("Fund").Preload("Fund.User").Preload("UserFund").Preload("UserDonate").Find(&transactions, "user_donate_id = ?", ID).Error

	return transactions, err
}

func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.First(&transaction, "id = ?", ID).Error

	return transaction, err
}

func (r *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Create(&transaction).Error

	return transaction, err
}

func (r *repository) GetOneTransaction(ID string) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("Fund").Preload("Fund.User").Preload("User").Preload("UserDonate").Preload("UserFund").First(&transaction, "id = ?", ID).Error

	return transaction, err
}

func (r *repository) UpdateTransaction(status string, ID string,Fund models.Fund,donate int) error {
	var transaction models.Transaction
	var fund models.Fund

	r.db.Preload("Fund").Preload("Fund.User").Preload("User").Preload("UserDonate").Preload("UserFund").First(&transaction, ID)
	r.db.Preload("User").First(&fund,Fund.ID)

	if status != transaction.Status && status == "success" {
		transaction.Status = status
		fund.Donated = Fund.Donated + donate 
	}

	r.db.Save(&fund)
	err := r.db.Save(&transaction).Error

	return err
}

func (r *repository) FindTransactionss(FundID int) ([]models.Transaction, error) {
	var transaction []models.Transaction
	err := r.db.Preload("Fund").Preload("Fund.User").Preload("UserFund").Preload("UserDonate").Where("fund_id = ? AND status = ?", FundID, "success").Find(&transaction).Error
	return transaction, err

}
func (r *repository) FindTransactionx(FundID int, UserFundID int) ([]models.Transaction, error) {
	var transaction []models.Transaction
	err := r.db.Preload("Fund").Preload("Fund.User").Preload("UserFund").Preload("UserDonate").Where("fund_id = ? AND user_fund_id = ?", FundID, UserFundID).Find(&transaction).Error
	return transaction, err

}

func (r *repository) FindTransactionall(FundID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("Fund").Preload("Fund.User").Preload("UserFund").Preload("UserDonate").Where("fund_id = ?", FundID).First(&transaction).Error

	return transaction, err
}


func (r *repository) GetFundTransaction(ID int) (models.Fund, error) {
	var fund models.Fund
	err := r.db.Preload("User").First(&fund, ID).Error

	return fund, err
}
