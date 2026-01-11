package routes

import (
	"build-in-public/internal/handlers"
	middleware "build-in-public/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	users := r.Group("/users")
	users.Use(middleware.RequireAuth())
	{
		users.GET("/me", handlers.Me)
	}

}
