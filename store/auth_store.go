package store

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/eternaleight/go-backend/models"
)

type AuthStore struct {
	DB *gorm.DB
}

func NewAuthStore(db *gorm.DB) *AuthStore {
	return &AuthStore{DB: db}
}

func (s *AuthStore) RegisterUser(username, email, password string) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}
	result := s.DB.Create(user)
	return user, result.Error
}

func (s *AuthStore) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := s.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (s *AuthStore) ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
