package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func loadConfig() string {
	// .envファイルを読み込む
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	dsn := fmt.Sprintf("host=%s port=%s user=postgres dbname=%s password=%s sslmode=disable",
		os.Getenv("HOST"),
		os.Getenv("PORT"),
		os.Getenv("DBNAME"),
		os.Getenv("PASSWORD"),
	)
	return dsn
}
