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
)

var db *gorm.DB

func main() {
	var err error

	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	// Read DATABASE_URL from Supabase
	dsn := fmt.Sprintf("host=%s port=%s user=postgres dbname=%s password=%s sslmode=disable",
		os.Getenv("HOST"),
		os.Getenv("PORT"),
		os.Getenv("DBNAME"),
		os.Getenv("PASSWORD"),
	)

	// Connect to the database using gorm
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Auto Migrate the User model
	db.AutoMigrate(&models.User{}, &models.Post{}, &models.Profile{}, &models.Product{}, &models.Purchase{})
	// Added Post and Profile models for migration as well

	r := gin.Default()
	r.RedirectTrailingSlash = false
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "Authorization") // 'Authorization' header allowed
	r.Use(cors.New(config))

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

	products := r.Group("/api/products").Use(middlewares.IsAuthenticated())
	{
		products.POST("/", handlers.CreateProduct)
		products.GET("/", handlers.ListProducts)
		products.GET("/:id", handlers.GetProduct)
		products.PUT("/:id", handlers.UpdateProduct)
		products.DELETE("/:id", handlers.DeleteProduct)
	}

	purchase := r.Group("/api/purchase").Use(middlewares.IsAuthenticated())
	{
		purchase.POST("/", handlers.PurchaseProduct)
	}

	r.Run(":8001")
}
