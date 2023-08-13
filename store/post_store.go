package store

import (
	"github.com/eternaleight/go-backend/models"
)

func CreatePost(post *models.Post) error {
	return DB.Create(post).Error
}

func GetLatestPosts() ([]models.Post, error) {
	var posts []models.Post
	if err := DB.Order("created_at desc").Limit(10).Preload("Author").Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}
