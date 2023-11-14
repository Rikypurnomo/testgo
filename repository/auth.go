package repository

import (
	"testgo/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(user models.User) error
	Login(email string) (models.User, error)
	CheckProfile(ID int) (models.User, error)
}

type authRepository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{
		DB: db,
	}
}

func (r *authRepository) Register(user models.User) error {
	err := r.DB.Create(&user).Error

	return err
}

func (r *authRepository) Login(email string) (models.User, error) {
	var user models.User
	err := r.DB.First(&user, "email=?", email).Error

	return user, err
}

func (r *authRepository) CheckProfile(ID int) (models.User, error) {
	var user models.User
	err := r.DB.First(&user, ID).Error

	return user, err
}
