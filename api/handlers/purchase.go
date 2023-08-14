package handlers

import (
	"github.com/eternaleight/go-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	"net/http"
)

// PurchaseProduct は商品の購入とStripeを使用した支払いを処理
func PurchaseProduct(c *gin.Context) {
	// Stripeのシークレットキーを設定
	stripe.Key = "YOUR_STRIPE_SECRET_KEY"

	var purchase models.Purchase
	if err := c.ShouldBindJSON(&purchase); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: データベースから関連する商品を取得
	var product models.Product
	// db.First(&product, purchase.ProductID)  // 仮のデータベースクエリ

	// TokenはStripe CheckoutまたはElementsを使用して作成されます
	token := c.PostForm("stripeToken")

	// ユーザーのカードを請求
	params := &stripe.ChargeParams{
		Amount:      stripe.Int64(int64(product.Price * 100)), // 価格をセントに変換
		Currency:    stripe.String("jpy"),                     // 通貨を日本円に変更
		Description: stripe.String("商品の請求"),
	}
	params.SetSource(token)
	_, err := charge.New(params)
	if err != nil {
		// エラーを処理
		c.JSON(http.StatusInternalServerError, gin.H{"error": "カードの請求に失敗しました"})
		return
	}

	// 支払いが成功した後、購入の詳細をデータベースに保存
	c.JSON(http.StatusOK, gin.H{"data": "商品は正常に購入されました"})
}
