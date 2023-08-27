package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/eternaleight/go-backend/models"
)

// 新しい商品を作成
func (h *Handler) CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "商品のデータの形式が正しくありません。"})
		return
	}

	err := h.ProductStore.CreateProduct(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "データベースに商品を保存できませんでした。"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// 全商品をリストとして取得
func (h *Handler) ListProducts(c *gin.Context) {
	products, err := h.ProductStore.ListProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "商品のリストの取得に失敗しました。"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": products})
}

// 指定されたIDの商品を取得
func (h *Handler) GetProductByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "商品IDが無効です。"})
		return
	}

	product, err := h.ProductStore.GetProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "商品の情報を取得できませんでした。"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": product})
}

// 指定されたIDの商品情報を更新
func (h *Handler) UpdateProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "商品IDが無効です。"})
		return
	}

	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "商品のデータの形式が正しくありません。"})
		return
	}

	err = h.ProductStore.UpdateProduct(uint(id), &product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "商品の更新に失敗しました。"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// 指定されたIDの商品を削除
func (h *Handler) DeleteProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "商品IDが無効です。"})
		return
	}

	err = h.ProductStore.DeleteProduct(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "商品の削除に失敗しました。"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "商品は正常に削除されました。"})
}
