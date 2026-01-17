package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"os"
	"time"

	"build-in-public/internal/config"
	"build-in-public/internal/dto"
	"build-in-public/internal/models"
	"build-in-public/internal/services"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

// generateStateToken generates a random state token for OAuth
func generateStateToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

// GoogleLogin godoc
// @Summary      Initiate Google OAuth login
// @Description  Redirects user to Google OAuth consent page
// @Tags         OAuth
// @Success      302
// @Router       /auth/google [get]
func GoogleLogin(c *gin.Context) {
	state := generateStateToken()

	// Store state in session/cookie for verification (in production, use Redis or DB)
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("oauth_state", state, 600, "/", "", false, true) // 10 minutes

	url := services.OAuth.Google.AuthCodeURL(state)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

// GoogleCallback godoc
// @Summary      Google OAuth callback
// @Description  Handles the callback from Google OAuth
// @Tags         OAuth
// @Param        code  query string true "Authorization code"
// @Param        state query string true "State token"
// @Success      200 {object} dto.SuccessResponse
// @Failure      400 {object} dto.ErrorResponse
// @Failure      401 {object} dto.ErrorResponse
// @Failure      500 {object} dto.ErrorResponse
// @Router       /auth/google/callback [get]
func GoogleCallback(c *gin.Context) {
	// Verify state
	state := c.Query("state")
	storedState, err := c.Cookie("oauth_state")
	if err != nil || state != storedState {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Error: "Invalid state token",
		})
		return
	}

	// Clear state cookie
	c.SetCookie("oauth_state", "", -1, "/", "", false, true)

	// Exchange code for token
	code := c.Query("code")
	token, err := services.OAuth.Google.Exchange(c.Request.Context(), code)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: "Failed to exchange token",
		})
		return
	}

	// Get user info from Google
	userInfo, err := services.GetGoogleUserInfo(c.Request.Context(), token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: "Failed to get user info",
		})
		return
	}

	// Handle user creation/login
	handleOAuthLogin(c, userInfo, token, models.OAuthGoogle)
}

// GitHubLogin godoc
// @Summary      Initiate GitHub OAuth login
// @Description  Redirects user to GitHub OAuth consent page
// @Tags         OAuth
// @Success      302
// @Router       /auth/github [get]
func GitHubLogin(c *gin.Context) {
	state := generateStateToken()

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("oauth_state", state, 600, "/", "", false, true)

	url := services.OAuth.GitHub.AuthCodeURL(state)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

// GitHubCallback godoc
// @Summary      GitHub OAuth callback
// @Description  Handles the callback from GitHub OAuth
// @Tags         OAuth
// @Param        code  query string true "Authorization code"
// @Param        state query string true "State token"
// @Success      200 {object} dto.SuccessResponse
// @Failure      400 {object} dto.ErrorResponse
// @Failure      401 {object} dto.ErrorResponse
// @Failure      500 {object} dto.ErrorResponse
// @Router       /auth/github/callback [get]
func GitHubCallback(c *gin.Context) {
	// Verify state
	state := c.Query("state")
	storedState, err := c.Cookie("oauth_state")
	if err != nil || state != storedState {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Error: "Invalid state token",
		})
		return
	}

	c.SetCookie("oauth_state", "", -1, "/", "", false, true)

	code := c.Query("code")
	token, err := services.OAuth.GitHub.Exchange(c.Request.Context(), code)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: "Failed to exchange token",
		})
		return
	}

	userInfo, err := services.GetGitHubUserInfo(c.Request.Context(), token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: "Failed to get user info",
		})
		return
	}

	handleOAuthLogin(c, userInfo, token, models.OAuthGithub)
}

// LinkedInLogin godoc
// @Summary      Initiate LinkedIn OAuth login
// @Description  Redirects user to LinkedIn OAuth consent page
// @Tags         OAuth
// @Success      302
// @Router       /auth/linkedin [get]
func LinkedInLogin(c *gin.Context) {
	state := generateStateToken()

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("oauth_state", state, 600, "/", "", false, true)

	url := services.OAuth.LinkedIn.AuthCodeURL(state)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

// LinkedInCallback godoc
// @Summary      LinkedIn OAuth callback
// @Description  Handles the callback from LinkedIn OAuth
// @Tags         OAuth
// @Param        code  query string true "Authorization code"
// @Param        state query string true "State token"
// @Success      200 {object} dto.SuccessResponse
// @Failure      400 {object} dto.ErrorResponse
// @Failure      401 {object} dto.ErrorResponse
// @Failure      500 {object} dto.ErrorResponse
// @Router       /auth/linkedin/callback [get]
func LinkedInCallback(c *gin.Context) {
	// Verify state
	state := c.Query("state")
	storedState, err := c.Cookie("oauth_state")
	if err != nil || state != storedState {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Error: "Invalid state token",
		})
		return
	}

	c.SetCookie("oauth_state", "", -1, "/", "", false, true)

	code := c.Query("code")
	token, err := services.OAuth.LinkedIn.Exchange(c.Request.Context(), code)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: "Failed to exchange token",
		})
		return
	}

	userInfo, err := services.GetLinkedInUserInfo(c.Request.Context(), token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: "Failed to get user info",
		})
		return
	}

	handleOAuthLogin(c, userInfo, token, models.OAuthLinkedIn)
}

// handleOAuthLogin handles the common OAuth login logic
func handleOAuthLogin(c *gin.Context, userInfo *services.OAuthUserInfo, token *oauth2.Token, provider models.OAuthProvider) {
	var user models.User
	var oauthAccount models.OAuthAccount

	// Check if OAuth account exists
	err := config.DB.Where("provider = ? AND provider_uid = ?", provider, userInfo.ID).
		First(&oauthAccount).Error

	if err == nil {
		// OAuth account exists, load the associated user
		if err := config.DB.First(&user, "id = ?", oauthAccount.UserID).Error; err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
				Error: "Failed to load user",
			})
			return
		}

		// Update OAuth account tokens
		oauthAccount.AccessToken = token.AccessToken
		oauthAccount.RefreshToken = token.RefreshToken
		if !token.Expiry.IsZero() {
			oauthAccount.ExpiresAt = &token.Expiry
		}
		oauthAccount.AvatarURL = userInfo.AvatarURL
		config.DB.Save(&oauthAccount)
	} else {
		// Check if user exists by email
		err = config.DB.Where("email = ?", userInfo.Email).First(&user).Error

		if err != nil {
			// Create new user
			user = models.User{
				FirstName:     userInfo.FirstName,
				LastName:      stringPtr(userInfo.LastName),
				Email:         userInfo.Email,
				EmailVerified: true, // OAuth emails are typically verified
				PhoneVerified: false,
			}

			if err := config.DB.Create(&user).Error; err != nil {
				c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
					Error: "Failed to create user",
				})
				return
			}
		}

		// Create OAuth account
		expiresAt := token.Expiry
		oauthAccount = models.OAuthAccount{
			UserID:       user.ID,
			Provider:     provider,
			ProviderUID:  userInfo.ID,
			Email:        userInfo.Email,
			AvatarURL:    userInfo.AvatarURL,
			AccessToken:  token.AccessToken,
			RefreshToken: token.RefreshToken,
		}
		if !expiresAt.IsZero() {
			oauthAccount.ExpiresAt = &expiresAt
		}

		if err := config.DB.Create(&oauthAccount).Error; err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
				Error: "Failed to link OAuth account",
			})
			return
		}
	}

	// Create session
	session := models.Session{
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	}

	if err := config.DB.Create(&session).Error; err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: "Failed to create session",
		})
		return
	}

	// Set session cookie
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

	// Redirect to frontend callback page
	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		frontendURL = "http://localhost:5173"
	}

	c.Redirect(http.StatusTemporaryRedirect, frontendURL+"/callback")
}

// stringPtr returns a pointer to the string
func stringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
