package register

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google" // Import Google OAuth2 endpoint
)

// ExportedOAuth2Config is a public OAuth2 configuration variable.
var ExportedOAuth2Config = &oauth2.Config{
	RedirectURL:  os.Getenv("REDIRECT_URL"),
	ClientID:     os.Getenv("CLIENT_ID"),
	ClientSecret: os.Getenv("CLIENT_SECRET"),
	Scopes:       []string{"profile", "email"},
	Endpoint:     google.Endpoint, // Set Google OAuth2 endpoint
}
