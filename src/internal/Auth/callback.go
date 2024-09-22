package Auth

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/tobbensol/elga_3_website/internal/Models/User"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

type DiscordUser struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

func Callback(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		if code == "" {
			http.Error(w, "No code in request", http.StatusBadRequest)
			return
		}

		cookie, err := r.Cookie("code_verifier")
		if err != nil {
			http.Error(w, "No code_verifier cookie found", http.StatusBadRequest)
			return
		}

		verifier := cookie.Value

		// delete the cookie
		http.SetCookie(w, &http.Cookie{
			Name:     "code_verifier",
			Value:    "",
			MaxAge:   -1,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
			Path:     "/auth/",
		})

		// Exchange the authorization code for an access token
		token, err := oauthConfig.Exchange(context.Background(), code, oauth2.VerifierOption(verifier))
		if err != nil {
			log.Printf("Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
			return
		}

		client := &http.Client{}

		request, err := http.NewRequest("GET", "https://discord.com/api/v10/users/@me", nil)
		if err != nil {
			log.Printf("unable to make a request for some reason: "+err.Error(), http.StatusInternalServerError)
			return
		}
		request.Header.Add("Authorization", "Bearer "+token.AccessToken)
		response, err := client.Do(request)
		if err != nil {
			log.Printf("Request error: %v", err)
			http.Error(w, "Unable to make request: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer response.Body.Close()

		// Check for HTTP status code
		if response.StatusCode != http.StatusOK {
			log.Printf("Discord API returned non-200 status: %d", response.StatusCode)
			http.Error(w, "Discord API returned an error", http.StatusInternalServerError)
			return
		}

		// Parse the JSON response
		var discordUser DiscordUser
		err = json.NewDecoder(response.Body).Decode(&discordUser)
		if err != nil {
			log.Printf("unable to make a request for some reason: "+err.Error(), http.StatusInternalServerError)
			return
		}

		//User.DeleteAll(db)

		// Look if user exists
		exists, err := User.DiscordUserExists(db, discordUser.ID)
		if !exists {
			// User does not exist, so make one
			User.Create(db, discordUser.Username, discordUser.ID, discordUser.Avatar, token.AccessToken, token.RefreshToken, token.Expiry)
		}

		fmt.Println(User.Get(db))

		// put the user ID in a cookie, THIS IS NOT SAFE AT ALL, BUT IT'S THE IDENTIFICATION FOR THE TIME BEING
		http.SetCookie(w, &http.Cookie{
			Name:     "authorization",
			Value:    discordUser.ID,
			MaxAge:   (int)(time.Hour.Seconds() * 24), // Cookie expiration time, 5 minutes for example
			HttpOnly: true,                            // Prevent JavaScript access to the cookie
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
			Path:     "/",
		})

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
