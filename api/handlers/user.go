package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/eternaleight/go-backend/models"
)

func (h *Handler) GetUser(c *gin.Context) {
	// isAuthenticatedミドルウェアで設定されたuserIDを取得
	userID, _ := c.Get("userID")

	var user models.User
	if err := h.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ユーザーが見つかりませんでした。"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"id":       user.ID,
			"email":    user.Email,
			"username": user.Username,
		},
	})
}
