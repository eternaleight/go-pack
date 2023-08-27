package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/eternaleight/go-backend/api/handlers"
	"github.com/eternaleight/go-backend/api/middlewares"
)

func setupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// トレーリングスラッシュへのリダイレクトを無効にする
	r.RedirectTrailingSlash = false

	config := cors.DefaultConfig()
	config.AllowCredentials = true

	allowedOrigins := os.Getenv("ALLOWED_ORIGINS") // 環境変数から読み取る
	if allowedOrigins == "" {
		allowedOrigins = "http://localhost:3000" // デフォルト値
	}
	config.AllowOrigins = []string{allowedOrigins} // フロントエンドのオリジンに合わせて変更

	r.Use(cors.New(config))

	// 'Authorization'ヘッダーを許可するためにヘッダーを追加
	config.AllowHeaders = append(config.AllowHeaders, "Authorization")

	// 新しいハンドラのインスタンスを作成し、データベースを渡す
	handlersNewDB := handlers.NewHandler(db)

	// auth
	auth := r.Group("/api/auth")
	{
		auth.POST("/register", handlersNewDB.Register)
		auth.POST("/login", handlersNewDB.Login)
	}

	// posts
	posts := r.Group("/api/posts").Use(middlewares.IsAuthenticated())
	{
		posts.POST("", handlersNewDB.CreatePost)
		posts.GET("", handlersNewDB.GetLatestPosts)
	}

	// user
	user := r.Group("/api/user").Use(middlewares.IsAuthenticated())
	user.GET("", handlersNewDB.GetUser)

	// products
	products := r.Group("/api/products").Use(middlewares.IsAuthenticated())
	{
		products.POST("", handlersNewDB.CreateProduct)
		products.GET("", handlersNewDB.ListProducts)
		products.GET("/:id", handlersNewDB.GetProductByID)
		products.PUT("/:id", handlersNewDB.UpdateProduct)
		products.DELETE("/:id", handlersNewDB.DeleteProduct)
	}

	// purchase
	purchase := r.Group("/api/purchase").Use(middlewares.IsAuthenticated())
	purchase.POST("", handlersNewDB.CreatePurchase)

	return r
}
