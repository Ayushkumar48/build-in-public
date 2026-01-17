package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/linkedin"
)

type OAuthConfig struct {
	Google   *oauth2.Config
	GitHub   *oauth2.Config
	LinkedIn *oauth2.Config
}

type OAuthUserInfo struct {
	ID        string
	Email     string
	FirstName string
	LastName  string
	AvatarURL string
	Provider  string
}

var OAuth *OAuthConfig

func InitOAuth() {
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}

	OAuth = &OAuthConfig{
		Google: &oauth2.Config{
			ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
			ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
			RedirectURL:  baseURL + "/auth/google/callback",
			Scopes: []string{
				"https://www.googleapis.com/auth/userinfo.email",
				"https://www.googleapis.com/auth/userinfo.profile",
			},
			Endpoint: google.Endpoint,
		},
		GitHub: &oauth2.Config{
			ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
			ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
			RedirectURL:  baseURL + "/auth/github/callback",
			Scopes: []string{
				"user:email",
				"read:user",
			},
			Endpoint: github.Endpoint,
		},
		LinkedIn: &oauth2.Config{
			ClientID:     os.Getenv("LINKEDIN_CLIENT_ID"),
			ClientSecret: os.Getenv("LINKEDIN_CLIENT_SECRET"),
			RedirectURL:  baseURL + "/auth/linkedin/callback",
			Scopes: []string{
				"openid",
				"profile",
				"email",
			},
			Endpoint: linkedin.Endpoint,
		},
	}
}

// GetGoogleUserInfo fetches user information from Google
func GetGoogleUserInfo(ctx context.Context, token *oauth2.Token) (*OAuthUserInfo, error) {
	client := OAuth.Google.Client(ctx, token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, fmt.Errorf("failed to get user info: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var googleUser struct {
		ID            string `json:"id"`
		Email         string `json:"email"`
		VerifiedEmail bool   `json:"verified_email"`
		Name          string `json:"name"`
		GivenName     string `json:"given_name"`
		FamilyName    string `json:"family_name"`
		Picture       string `json:"picture"`
	}

	if err := json.Unmarshal(data, &googleUser); err != nil {
		return nil, fmt.Errorf("failed to unmarshal user info: %w", err)
	}

	return &OAuthUserInfo{
		ID:        googleUser.ID,
		Email:     googleUser.Email,
		FirstName: googleUser.GivenName,
		LastName:  googleUser.FamilyName,
		AvatarURL: googleUser.Picture,
		Provider:  "google",
	}, nil
}

// GetGitHubUserInfo fetches user information from GitHub
func GetGitHubUserInfo(ctx context.Context, token *oauth2.Token) (*OAuthUserInfo, error) {
	client := OAuth.GitHub.Client(ctx, token)

	// Get user info
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		return nil, fmt.Errorf("failed to get user info: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var githubUser struct {
		ID        int64  `json:"id"`
		Login     string `json:"login"`
		Name      string `json:"name"`
		Email     string `json:"email"`
		AvatarURL string `json:"avatar_url"`
	}

	if err := json.Unmarshal(data, &githubUser); err != nil {
		return nil, fmt.Errorf("failed to unmarshal user info: %w", err)
	}

	// If email is not public, fetch from emails endpoint
	email := githubUser.Email
	if email == "" {
		emailResp, err := client.Get("https://api.github.com/user/emails")
		if err == nil {
			defer emailResp.Body.Close()
			emailData, err := io.ReadAll(emailResp.Body)
			if err == nil {
				var emails []struct {
					Email      string `json:"email"`
					Primary    bool   `json:"primary"`
					Verified   bool   `json:"verified"`
					Visibility string `json:"visibility"`
				}
				if json.Unmarshal(emailData, &emails) == nil {
					for _, e := range emails {
						if e.Primary && e.Verified {
							email = e.Email
							break
						}
					}
					if email == "" && len(emails) > 0 && emails[0].Verified {
						email = emails[0].Email
					}
				}
			}
		}
	}

	// Parse name into first and last
	firstName := githubUser.Name
	lastName := ""
	if githubUser.Name != "" {
		names := splitName(githubUser.Name)
		firstName = names[0]
		if len(names) > 1 {
			lastName = names[1]
		}
	} else {
		firstName = githubUser.Login
	}

	return &OAuthUserInfo{
		ID:        fmt.Sprintf("%d", githubUser.ID),
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		AvatarURL: githubUser.AvatarURL,
		Provider:  "github",
	}, nil
}

// GetLinkedInUserInfo fetches user information from LinkedIn
func GetLinkedInUserInfo(ctx context.Context, token *oauth2.Token) (*OAuthUserInfo, error) {
	client := OAuth.LinkedIn.Client(ctx, token)

	// Get user profile
	resp, err := client.Get("https://api.linkedin.com/v2/userinfo")
	if err != nil {
		return nil, fmt.Errorf("failed to get user info: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var linkedinUser struct {
		Sub           string `json:"sub"`
		Email         string `json:"email"`
		EmailVerified bool   `json:"email_verified"`
		GivenName     string `json:"given_name"`
		FamilyName    string `json:"family_name"`
		Name          string `json:"name"`
		Picture       string `json:"picture"`
	}

	if err := json.Unmarshal(data, &linkedinUser); err != nil {
		return nil, fmt.Errorf("failed to unmarshal user info: %w", err)
	}

	return &OAuthUserInfo{
		ID:        linkedinUser.Sub,
		Email:     linkedinUser.Email,
		FirstName: linkedinUser.GivenName,
		LastName:  linkedinUser.FamilyName,
		AvatarURL: linkedinUser.Picture,
		Provider:  "linkedin",
	}, nil
}

// splitName splits a full name into first and last name
func splitName(fullName string) []string {
	var firstName, lastName string
	names := []string{}

	// Simple split by space
	for i, char := range fullName {
		if char == ' ' {
			firstName = fullName[:i]
			lastName = fullName[i+1:]
			break
		}
	}

	if firstName == "" {
		firstName = fullName
	}

	names = append(names, firstName)
	if lastName != "" {
		names = append(names, lastName)
	}

	return names
}
