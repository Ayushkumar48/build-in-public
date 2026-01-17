package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"build-in-public/internal/config"
	"build-in-public/internal/routes"
	"build-in-public/internal/services"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ No .env file found, using system env")
	}

	config.ConnectDatabase()
	services.InitOAuth()

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // SvelteKit dev
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Cookie"},
		AllowCredentials: true, // 🔥 REQUIRED FOR COOKIES
		ExposeHeaders:    []string{"Set-Cookie"},
		MaxAge:           12 * time.Hour,
	}))

	routes.RegisterAuthRoutes(r)
	routes.RegisterUserRoutes(r)
	r.Run(":" + os.Getenv("APP_PORT"))
}
