package dto

import (
	"build-in-public/internal/models"
	"time"

	"github.com/google/uuid"
)

type SocialAccountResponse struct {
	Platform string `json:"platform"`
	Username string `json:"username"`
	URL      string `json:"url"`
}

type OAuthProviderResponse struct {
	Provider string `json:"provider"`
}

type UserResponse struct {
	ID              uuid.UUID               `json:"id"`
	FirstName       string                  `json:"first_name"`
	LastName        string                  `json:"last_name"`
	Email           string                  `json:"email"`
	Username        string                  `json:"username"`
	Phone           string                  `json:"phone"`
	Socials         []SocialAccountResponse `json:"socials"`
	OAuthProviders  []OAuthProviderResponse `json:"oauth_providers"`
	CreatedAt       time.Time               `json:"createdAt"`
	UpdatedAt       time.Time               `json:"updatedAt"`
	EmailVerified   bool                    `json:"email_verified"`
	PhoneNoVerified bool                    `json:"phone_no_verified"`
	Gender          models.Gender           `json:"gender"`
	DateOfBirth     time.Time               `json:"date_of_birth"`
	City            string                  `json:"city"`
	Bio             string                  `json:"bio"`
}

func ToUserResponse(user models.User) UserResponse {
	socials := make([]SocialAccountResponse, 0, len(user.Socials))
	for _, s := range user.Socials {
		socials = append(socials, SocialAccountResponse{
			Platform: string(s.Platform),
			Username: s.Username,
			URL:      s.URL,
		})
	}

	oauthProviders := make([]OAuthProviderResponse, 0, len(user.OAuthAccounts))
	for _, acc := range user.OAuthAccounts {
		oauthProviders = append(oauthProviders, OAuthProviderResponse{
			Provider: string(acc.Provider),
		})
	}

	return UserResponse{
		ID:              user.ID,
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		Email:           user.Email,
		Username:        user.Username,
		Phone:           user.Phone,
		EmailVerified:   user.EmailVerified,
		PhoneNoVerified: user.PhoneVerified,
		Socials:         socials,
		OAuthProviders:  oauthProviders,
		Gender:          user.Gender,
		DateOfBirth:     user.DateOfBirth,
		City:            user.City,
		Bio:             user.Bio,
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
	}
}
