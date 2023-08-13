package store

import (
	"github.com/eternaleight/go-backend/models"
)

func CreateUser(user *models.User) error {
	return DB.Create(user).Error
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
