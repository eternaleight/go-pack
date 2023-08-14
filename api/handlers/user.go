package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/eternaleight/go-backend/store"
)


// ユーザー情報を取得
func (h *Handler) GetUser(c *gin.Context) {
	// ミドルウェアからuserIDを取得
	userID := c.MustGet("userID").(int)

	// UserStoreのインスタンスを生成
	userStore := store.NewUserStore(h.DB)
	// データベースからユーザー情報を取得
	user, err := userStore.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found or database error"})
		return
	}

	// ユーザー情報をレスポンスとして返す
	c.JSON(http.StatusOK, gin.H{"user": user})
}
