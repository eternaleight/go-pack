package handlers

import (
	"github.com/eternaleight/go-backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreatePost(c *gin.Context) {
	var input struct {
		Content string `json:"content"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "投稿内容がありません"})
		return
	}

	// isAuthenticatedミドルウェアで設定されたuserIDを取得
	userID, _ := c.Get("userID")

	post := models.Post{
		Content:  input.Content,
		AuthorID: userID.(int), // userIDの型変換
	}
	if err := db.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "サーバーエラーです。"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"post": post})
}

func GetLatestPosts(c *gin.Context) {
	var posts []models.Post

	if err := db.Order("created_at desc").Limit(10).Preload("Author").Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "サーバーエラーです。"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"posts": posts})
}
