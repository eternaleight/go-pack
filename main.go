package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/eternaleight/go-backend/api/handlers"
	"github.com/eternaleight/go-backend/api/middlewares"
	"github.com/eternaleight/go-backend/models"
	"github.com/eternaleight/go-backend/store"
)

var db *gorm.DB

func main() {
	var err error

	// .envファイルを読み込む
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	// SupabaseからDATABASE_URLを読み込む
	dsn := fmt.Sprintf("host=%s port=%s user=postgres dbname=%s password=%s sslmode=disable",
		os.Getenv("HOST"),
		os.Getenv("PORT"),
		os.Getenv("DBNAME"),
		os.Getenv("PASSWORD"),
	)

	// gormを使ってデータベースに接続
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// ユーザーモデルを自動マイグレーション
	db.AutoMigrate(&models.User{}, &models.Post{}, &models.Profile{}, &models.Product{}, &models.Purchase{})

	r := gin.Default()
  // トレーリングスラッシュへのリダイレクトを無効にする
	r.RedirectTrailingSlash = false
	config := cors.DefaultConfig()
	r.Use(cors.New(config))
	config.AllowAllOrigins = true
	// 'Authorization'ヘッダーを許可するためにヘッダーを追加
	config.AllowHeaders = append(config.AllowHeaders, "Authorization")

	// 新しいハンドラのインスタンスを作成し、データベースを渡す
	handlersNewDB := handlers.NewHandler(db)

	auth := r.Group("/api/auth")
	{
		auth.POST("/register", handlersNewDB.Register)
		auth.POST("/login", handlersNewDB.Login)
	}

	posts := r.Group("/api/posts").Use(middlewares.IsAuthenticated())
	{
		posts.POST("/", handlersNewDB.CreatePost)
		posts.GET("/", handlersNewDB.GetLatestPosts)
	}

	user := r.Group("/api/user").Use(middlewares.IsAuthenticated())
	{
		user.GET("/", handlersNewDB.GetUser)
	}

	// ストアのインスタンスを作成
	productStore := store.NewProductStore(db)
	purchaseStore := store.NewPurchaseStore(db)

	// ハンドラのインスタンスを作成
	productHandler := handlers.NewProductHandler(productStore)
	purchaseHandler := handlers.NewPurchaseHandler(purchaseStore)

	products := r.Group("/api/products").Use(middlewares.IsAuthenticated())
	{
		products.POST("/", productHandler.CreateProduct)
		products.GET("/", productHandler.ListProducts)
		products.GET("/:id", productHandler.GetProductByID)
		products.PUT("/:id", productHandler.UpdateProduct)
		products.DELETE("/:id", productHandler.DeleteProduct)
	}

	purchase := r.Group("/api/purchase").Use(middlewares.IsAuthenticated())
	{
		purchase.POST("/", purchaseHandler.CreatePurchase)
	}

	r.Run(":8001")
}
