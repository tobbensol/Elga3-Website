package Auth

import (
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

func getConfig() *oauth2.Config {
	envFile, _ := godotenv.Read("./.env")
	config := &oauth2.Config{
		ClientID:     envFile["CLIENT_ID"],
		ClientSecret: envFile["CLIENT_SECRET"],
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://discord.com/oauth2/authorize",
			TokenURL: "https://discord.com/api/oauth2/token",
		},
		Scopes:      []string{"identify"},
		RedirectURL: "http://localhost:8000/auth/callback",
	}
	return config
}
