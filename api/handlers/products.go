package handlers

import (
	"github.com/eternaleight/go-backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateProduct は新しい商品を作成
func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// TODO: データベースに商品を保存
	c.JSON(http.StatusOK, gin.H{"data": product})
}

// ListProducts は利用可能な商品をリストアップ
func ListProducts(c *gin.Context) {
	var products []models.Product
	// TODO: データベースから商品を取得
	c.JSON(http.StatusOK, gin.H{"data": products})
}

// GetProduct はIDに基づいて商品の詳細を取得
func GetProduct(c *gin.Context) {
	var product models.Product
	// id := c.Param("id") // 現段階では未使用のためコメントアウト
	// TODO: データベースから指定されたIDの商品を取得
	c.JSON(http.StatusOK, gin.H{"data": product})
}

// UpdateProduct は既存の商品の詳細を更新
func UpdateProduct(c *gin.Context) {
	var product models.Product
	// id := c.Param("id") // 現段階では未使用のためコメントアウト
	// TODO: データベースで指定されたIDの商品を更新
	c.JSON(http.StatusOK, gin.H{"data": product})
}

// DeleteProduct はIDに基づいて商品を削除
func DeleteProduct(c *gin.Context) {
	// id := c.Param("id") // 現段階では未使用のためコメントアウト
	// TODO: データベースから指定されたIDの商品を削除
	c.JSON(http.StatusOK, gin.H{"data": "商品は正常に削除されました"})
}
