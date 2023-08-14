package handlers

import (
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/eternaleight/go-backend/store"
)

// DBのインスタンスを持つ構造体
type Handler struct {
	DB *gorm.DB
}

// 新しいHandlerのインスタンスを初期化
func NewHandler(db *gorm.DB) *Handler {
	return &Handler{DB: db}
}

// 新しいユーザーを登録
func (h *Handler) Register(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// リクエストからJSONデータをバインド
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// AuthStoreのインスタンスを生成
	authStore := store.NewAuthStore(h.DB)

	// ユーザー登録情報をデータベースに保存
	user, err := authStore.RegisterUser(input.Username, input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// 登録されたユーザー情報をレスポンスとして返す
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// ユーザーのログインを処理
func (h *Handler) Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// リクエストからJSONデータをバインド
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// AuthStoreのインスタンスを生成
	authStore := store.NewAuthStore(h.DB)

	// メールアドレスを使ってユーザーを取得
	user, err := authStore.GetUserByEmail(input.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "メールアドレスが存在しません"})
		return
	}

	// ユーザーのパスワードを検証
	err = authStore.ComparePassword(user.Password, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "パスワードが間違っています"})
		return
	}

	// JWTトークンを生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "トークンの生成に失敗しました"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
