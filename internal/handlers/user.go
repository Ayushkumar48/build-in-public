package handlers

import (
	"build-in-public/internal/dto"
	"build-in-public/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Me godoc
// @Summary      Get current user
// @Description  Returns logged-in user
// @Tags         Auth
// @Success      200 {object} dto.UserResponse
// @Failure      401 {object} dto.ErrorResponse
// @Router       /users/me [get]
func Me(c *gin.Context) {
	userAny, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Error: "unauthorized",
		})
		return
	}

	user, ok := userAny.(models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "invalid user type"})
		return
	}

	response := dto.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		PhoneNo:   user.PhoneNo,
		LinkedIn:  user.LinkedIn,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}
