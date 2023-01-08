package repositories

import (
	"dewetour/models"

	"gorm.io/gorm"
)

type TripRepository interface {
	FindTrip() ([]models.Trip, error)
	GetTrip(ID int) (models.Trip, error)
	CreateTrip(user models.Trip) (models.Trip, error)
	UpdateTrip(user models.Trip) (models.Trip, error)
	DeleteTrip(user models.Trip, ID int) (models.Trip, error)
}

type repositoryTrip struct {
	db *gorm.DB
}

func RepositoryTrip(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTrip() ([]models.Trip, error) {
	var trips []models.Trip
	err := r.db.Preload("Country").Preload("User").Find(&trips).Error

	return trips, err
}
func (r *repository) GetTrip(ID int) (models.Trip, error) {
	var trip models.Trip
	err := r.db.Preload("Country").Preload("User").First(&trip, ID).Error

	return trip, err
}
func (r *repository) CreateTrip(trip models.Trip) (models.Trip, error) {
	// var user models.User
	err := r.db.Create(&trip).Error

	return trip, err
}

func (r *repository) UpdateTrip(trip models.Trip) (models.Trip, error) {
	// var user models.User
	err := r.db.Model(&trip).Updates(trip).Error

	return trip, err
}
func (r *repository) DeleteTrip(trip models.Trip, ID int) (models.Trip, error) {
	// var user models.User
	err := r.db.Preload("Country").Delete(&trip).Error

	return trip, err
}
