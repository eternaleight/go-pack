package store

import (
	"github.com/eternaleight/go-backend/models"
	"gorm.io/gorm"
)

type PostStore struct {
	DB *gorm.DB
}

func NewPostStore(db *gorm.DB) *PostStore {
	return &PostStore{DB: db}
}

func (s *PostStore) CreatePost(post *models.Post) error {
	return s.DB.Create(post).Error
}

func (s *PostStore) GetLatestPosts() ([]models.Post, error) {
	var posts []models.Post
	if err := s.DB.Order("created_at desc").Limit(10).Preload("Author").Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}
