package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Config struct {
	GoogleLoginConfig   oauth2.Config
	FacebookLoginConfig oauth2.Config
}

var AppConfig Config

const OauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func LoadConfig() {
	// Oauth configuration for Google
	AppConfig.GoogleLoginConfig = oauth2.Config{
		ClientID:     "117618649798-id6vcol34j1aj6653nkns31l1os9dssr.apps.googleusercontent.com",
		ClientSecret: "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",
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://kase.dhasatra.com/auth/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
	}

}
