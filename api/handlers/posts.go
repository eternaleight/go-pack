package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/eternaleight/go-backend/models"
)

func (h *Handler) CreatePost(c *gin.Context) {
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
	userIDValue, exists := c.Get("userID")
	fmt.Println("UserID retrieved in handler:", userIDValue) // この行を追加
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ユーザーIDが見つかりません"})
		return
	}

	// userIDの型確認
	userID, ok := userIDValue.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ユーザーIDの型が正しくありません"})
		return
	}

	post := models.Post{
		Content:  input.Content,
		AuthorID: userID,
	}
	if err := h.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "サーバーエラーです。"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"post": post})
}

func (h *Handler) GetLatestPosts(c *gin.Context) {
	var posts []models.Post

	if err := h.DB.Order("created_at desc").Limit(10).Preload("Author").Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "サーバーエラーです。"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"posts": posts})
}
