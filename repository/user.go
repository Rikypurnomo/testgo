package repository

import (
	"testgo/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	ListUsers() ([]models.User, error)
	GetUser(ID int) (models.User, error)
}

type userRepository struct {
	DB *gorm.DB
}

func NewUsersRepository(DB *gorm.DB) *userRepository {
	return &userRepository{
		DB: DB,
	}
}

func (r *userRepository) ListUsers() ([]models.User, error) {
	var users []models.User
	err := r.DB.Find(&users).Error

	return users, err
}

func (r *userRepository) GetUser(ID int) (models.User, error) {
	var user models.User
	err := r.DB.First(&user, ID).Error

	return user, err
}
