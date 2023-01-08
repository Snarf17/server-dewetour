package repositories

import (
	"dewetour/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindTransaction() ([]models.Transaction, error)
	GetTransaction(ID int) (models.Transaction, error)
	GetOneTransaction(ID string) (models.Transaction, error)
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
	UpdateTransaction(status string, ID int) (models.Transaction, error)
	// DeleteUsers(user models.User, ID int) (models.User, error)
}

type repositoryTransaction struct {
	db *gorm.DB
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTransaction() ([]models.Transaction, error) {
	var transaction []models.Transaction
	err := r.db.Preload("Trip").Preload("Trip.Country").Preload("User").Find(&transaction).Error

	return transaction, err
}

func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
	var Transaction models.Transaction
	err := r.db.Preload("Trip").Preload("Trip.Country").Preload("User").First(&Transaction, ID).Error

	return Transaction, err
}

//
func (r *repository) GetOneTransaction(ID string) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("Trip").Preload("Trip.Country").Preload("User").First(&transaction, "id = ?", ID).Error

	return transaction, err
}

func (r *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	// var user models.User
	err := r.db.Preload("Trip").Preload("Trip.Country").Preload("User").Create(&transaction).Error
	return transaction, err
}

func (r *repository) UpdateTransaction(status string, ID int) (models.Transaction, error) {
	var transaction models.Transaction
	r.db.Preload("trip").First(&transaction, ID)

	// If is different & Status is "success" decrement product quantity
	if status != transaction.Status && status == "success" {
		var trip models.Trip
		r.db.First(&trip, transaction.Trip.ID)
		trip.Quota = trip.Quota - 1
		r.db.Save(&trip)
	}

	transaction.Status = status

	err := r.db.Save(&transaction).Error

	return transaction, err
}

// func (r *repository) DeleteTransaction(Transaction models.Transaction, ID int) (models.Transaction, error) {
// 	// var user models.User
// 	err := r.db.Preload("Country").Delete(&Transaction).Error

// 	return Transaction, err
// }
