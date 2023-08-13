package main

import (
	"github.com/eternaleight/go-backend/api/handlers"
	"github.com/eternaleight/go-backend/api/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "Authorization") // 'Authorization' ヘッダーを許可
	r.Use(cors.New(config))

	auth := r.Group("/api/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
	}

	posts := r.Group("/api/posts").Use(middlewares.IsAuthenticated())
	{
		posts.POST("/", handlers.CreatePost)
		posts.GET("/", handlers.GetLatestPosts)
	}

	user := r.Group("/api/user").Use(middlewares.IsAuthenticated())
	{
		user.GET("/", handlers.GetUser)
	}

	r.Run(":8001")
}
