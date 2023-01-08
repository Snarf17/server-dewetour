package repositories

import (
	"dewetour/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUsers() ([]models.User, error)
	GetUsers(ID int) (models.User, error)
	CreateUsers(user models.User) (models.User, error)
	UpdateUsers(user models.User) (models.User, error)
	DeleteUsers(user models.User, ID int) (models.User, error)
}

type repository struct {
	db *gorm.DB
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error

	return users, err
}

func (r *repository) GetUsers(ID int) (models.User, error) {
	var user models.User
	err := r.db.First(&user, ID).Error

	return user, err
}
func (r *repository) CreateUsers(user models.User) (models.User, error) {
	// var user models.User
	err := r.db.Create(&user).Error

	return user, err
}
func (r *repository) UpdateUsers(user models.User) (models.User, error) {
	// var user models.User
	err := r.db.Save(&user).Error

	return user, err
}
func (r *repository) DeleteUsers(user models.User, ID int) (models.User, error) {
	// var user models.User
	err := r.db.Delete(&user).Error

	return user, err
}
