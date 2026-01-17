package routes

import (
	"build-in-public/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		// Traditional auth
		auth.POST("/signup", handlers.Signup)
		auth.POST("/login", handlers.Login)
		auth.POST("/logout", handlers.Logout)

		// OAuth - Google
		auth.GET("/google", handlers.GoogleLogin)
		auth.GET("/google/callback", handlers.GoogleCallback)

		// OAuth - GitHub
		auth.GET("/github", handlers.GitHubLogin)
		auth.GET("/github/callback", handlers.GitHubCallback)

		// OAuth - LinkedIn
		auth.GET("/linkedin", handlers.LinkedInLogin)
		auth.GET("/linkedin/callback", handlers.LinkedInCallback)
	}
}
