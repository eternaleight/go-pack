package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/eternaleight/go-backend/models"
)

// 新しい購入を作成するためのハンドラ
func (h *Handler) CreatePurchase(c *gin.Context) {
	var purchase models.Purchase

	// 購入データのJSONをパース
	if err := c.ShouldBindJSON(&purchase); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "購入データの形式が正しくない"})
		return
	}

	// 購入データをデータベースに保存
	err := h.PurchaseStore.CreatePurchase(&purchase)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "データベースに購入情報を保存できなかった"})
		return
	}

	// 保存に成功した場合のレスポンスを返す
	c.JSON(http.StatusOK, gin.H{"data": "商品の購入が成功した"})
}
