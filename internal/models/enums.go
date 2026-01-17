package models

type SocialPlatform string
type Gender string
type OAuthProvider string

const (
	PlatformGithub   SocialPlatform = "github"
	PlatformLinkedIn SocialPlatform = "linkedin"
	PlatformTwitter  SocialPlatform = "twitter"
	PlatformWebsite  SocialPlatform = "website"
)

const (
	GenderMale   Gender = "male"
	GenderFemale Gender = "female"
	GenderOther  Gender = "other"
)

const (
	OAuthGoogle    OAuthProvider = "google"
	OAuthGithub    OAuthProvider = "github"
	OAuthLinkedIn  OAuthProvider = "linkedin"
	OAuthMicrosoft OAuthProvider = "microsoft"
)
