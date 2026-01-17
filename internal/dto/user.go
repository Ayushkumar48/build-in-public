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

type CollegeResponse struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Domain  string    `json:"domain"`
	City    string    `json:"city,omitempty"`
	State   string    `json:"state,omitempty"`
	Country string    `json:"country,omitempty"`
}

type UserResponse struct {
	ID              uuid.UUID               `json:"id"`
	FirstName       string                  `json:"first_name"`
	LastName        *string                 `json:"last_name,omitempty"`
	Email           string                  `json:"email"`
	Username        *string                 `json:"username,omitempty"`
	Phone           *string                 `json:"phone,omitempty"`
	EmailVerified   bool                    `json:"email_verified"`
	PhoneNoVerified bool                    `json:"phone_no_verified"`
	Socials         []SocialAccountResponse `json:"socials"`
	OAuthProviders  []OAuthProviderResponse `json:"oauth_providers"`
	Gender          models.Gender           `json:"gender"`
	DateOfBirth     *time.Time              `json:"date_of_birth,omitempty"`
	City            *string                 `json:"city,omitempty"`
	Bio             *string                 `json:"bio,omitempty"`
	College         *CollegeResponse        `json:"college,omitempty"`
	CreatedAt       time.Time               `json:"created_at"`
	UpdatedAt       time.Time               `json:"updated_at"`
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

	var college *CollegeResponse
	if user.College != nil {
		college = &CollegeResponse{
			ID:      user.College.ID,
			Name:    user.College.Name,
			Domain:  user.College.Domain,
			City:    user.College.City,
			State:   user.College.State,
			Country: user.College.Country,
		}
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
		College:         college,
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
	}
}
