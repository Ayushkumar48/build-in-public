package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type College struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name      string    `gorm:"size:255;not null" json:"name"`
	Domain    string    `gorm:"size:255;not null;uniqueIndex" json:"domain"`
	City      string    `gorm:"size:255" json:"city"`
	State     string    `gorm:"size:255" json:"state"`
	Country   string    `gorm:"size:255" json:"country"`
	Users     []User    `gorm:"foreignKey:CollegeID" json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type SocialAccount struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID    uuid.UUID      `gorm:"type:uuid;not null;uniqueIndex:idx_user_platform"`
	Platform  SocialPlatform `gorm:"type:varchar(50);not null;uniqueIndex:idx_user_platform"`
	Username  string         `gorm:"size:255;not null" json:"username"`
	URL       string         `gorm:"size:255" json:"url"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type OAuthAccount struct {
	ID           uuid.UUID     `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID       uuid.UUID     `gorm:"type:uuid;not null;index"`
	Provider     OAuthProvider `gorm:"type:varchar(50);not null;uniqueIndex:idx_provider_uid"`
	ProviderUID  string        `gorm:"size:255;not null;uniqueIndex:idx_provider_uid"`
	Email        string        `gorm:"size:255"`
	AvatarURL    string        `gorm:"size:500"`
	AccessToken  string        `gorm:"type:text"`
	RefreshToken string        `gorm:"type:text"`
	ExpiresAt    *time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (OAuthAccount) TableName() string {
	return "oauth_accounts"
}

type User struct {
	ID            uuid.UUID       `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	FirstName     string          `gorm:"size:255;not null" json:"first_name"`
	LastName      *string         `gorm:"size:255;" json:"last_name"`
	Email         string          `gorm:"uniqueIndex;not null" json:"email"`
	EmailVerified bool            `gorm:"not null" json:"email_verified"`
	Username      *string         `gorm:"size:255;" json:"username"`
	Phone         *string         `gorm:"size:20" json:"phone"`
	PhoneVerified bool            `gorm:"not null" json:"phone_verified"`
	Gender        Gender          `gorm:"type:varchar(10);not null" json:"gender"`
	DateOfBirth   *time.Time      `gorm:"type:date" json:"date_of_birth"`
	City          *string         `gorm:"size:255" json:"city"`
	Bio           *string         `gorm:"size:255" json:"bio"`
	Password      *string         `json:"-"`
	OAuthAccounts []OAuthAccount  `gorm:"foreignKey:UserID" json:"oauth_accounts"`
	Socials       []SocialAccount `gorm:"foreignKey:UserID" json:"socials"`
	College       *College        `gorm:"foreignKey:CollegeID" json:"college"`
	CollegeID     *uuid.UUID      `gorm:"type:uuid" json:"college_id"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
	DeletedAt     gorm.DeletedAt  `gorm:"index" json:"-"`
}
