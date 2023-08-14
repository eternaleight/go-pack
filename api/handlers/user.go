package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/eternaleight/go-backend/store"
)

func (h *Handler) GetUser(c *gin.Context) {
	userID := c.MustGet("userID").(int)

	userStore := store.NewUserStore(h.DB)
	user, err := userStore.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found or database error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
