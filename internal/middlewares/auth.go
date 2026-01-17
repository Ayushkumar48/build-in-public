package middleware

import (
	"net/http"
	"time"

	"build-in-public/internal/config"
	"build-in-public/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionID, err := c.Cookie("session_id")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthenticated"})
			return
		}

		id, err := uuid.Parse(sessionID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid session"})
			return
		}

		var session models.Session
		if err := config.DB.First(&session, "id = ?", id).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "session not found"})
			return
		}

		if session.ExpiresAt.Before(time.Now()) {
			config.DB.Delete(&session)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "session expired"})
			return
		}

		var user models.User
		if err := config.DB.
			Preload("Socials", "deleted_at IS NULL").
			First(&user, "id = ?", session.UserID).Error; err != nil {

			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
			return
		}

		// Attach user to request context
		c.Set("user", user)

		c.Next()
	}
}
