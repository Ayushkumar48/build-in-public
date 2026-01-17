package handlers

import (
	"net/http"
	"time"

	"build-in-public/internal/config"
	"build-in-public/internal/dto"
	"build-in-public/internal/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type SignupRequest struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=8"`
	Phone     string `json:"phoneNo" binding:"max=20"`
	LinkedIn  string `json:"linkedIn" binding:"max=255"`
}
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// Signup godoc
// @Summary      Create a new user account
// @Description  Register a new user with email and password
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body SignupRequest true "Signup Request"
// @Success      201 {object} dto.UserResponse
// @Failure      400 {object} dto.ErrorResponse
// @Failure      409 {object} dto.ErrorResponse
// @Failure      500 {object} dto.ErrorResponse
// @Router       /auth/signup [post]
func Signup(c *gin.Context) {
	var req SignupRequest

	// 1. Validate input

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	// 2. Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: "Failed to hash password",
		})
		return
	}

	// 3. Create user
	hashed := string(hashedPassword)

	user := models.User{
		FirstName:     req.FirstName,
		LastName:      &req.LastName,
		Email:         req.Email,
		Password:      &hashed,
		Phone:         nil,
		EmailVerified: false,
		PhoneVerified: false,
	}
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusConflict, dto.ErrorResponse{
			Error: "Email already exists",
		})
		return
	}
	// Create session
	session := models.Session{
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	}

	if err := config.DB.Create(&session).Error; err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: "Failed to create a session",
		})
		return
	}

	// Set cookie with SameSite=Lax for localhost compatibility
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(
		"session_id",
		session.ID.String(),
		60*60*24*7, // 7 days
		"/",
		"",
		false, // true in production (HTTPS)
		true,  // HttpOnly
	)
	c.JSON(http.StatusCreated, dto.SuccessResponse{Success: "User created successfully"})
}

// Login godoc
// @Summary      Login to user account
// @Description  Authenticate user with email and password
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body LoginRequest true "Login Request"
// @Success      200 {object} dto.SuccessResponse
// @Failure      400 {object} dto.ErrorResponse
// @Failure      401 {object} dto.ErrorResponse
// @Failure      500 {object} dto.ErrorResponse
// @Router       /auth/login [post]
func Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Error: "Invalid Credentials",
		})
		return
	}

	if user.Password == nil {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Error: "Use OAuth login",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(*user.Password),
		[]byte(req.Password),
	); err != nil {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Error: "Invalid Credentials",
		})
		return
	}

	// Create session
	session := models.Session{
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	}

	if err := config.DB.Create(&session).Error; err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: "Failed to create a session",
		})
		return
	}

	// Set cookie with SameSite=Lax for localhost compatibility
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(
		"session_id",
		session.ID.String(),
		60*60*24*7, // 7 days
		"/",
		"",
		false, // true in production (HTTPS)
		true,  // HttpOnly
	)

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: "Logged In",
	})
}

// Logout godoc
// @Summary      Logout from user account
// @Description  Invalidate user session and clear cookies
// @Tags         Auth
// @Success      200 {object} dto.SuccessResponse
// @Router       /auth/logout [post]
func Logout(c *gin.Context) {
	sessionID, _ := c.Cookie("session_id")
	config.DB.Delete(&models.Session{}, "id = ?", sessionID)

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("session_id", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: "Logged Out",
	})
}
