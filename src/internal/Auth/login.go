package Auth

import (
	"golang.org/x/oauth2"
	"net/http"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
	codeVerifier := oauth2.GenerateVerifier()

	// Set the code_verifier in a cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "code_verifier",
		Value:    codeVerifier,
		MaxAge:   (int)(5 * time.Minute), // Cookie expiration time, 5 minutes for example
		HttpOnly: true,                   // Prevent JavaScript access to the cookie
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		Path:     "/auth",
	})
	// access type online = access even if the user is away
	url := oauthConfig.AuthCodeURL(
		"random_state", // this should idealy be random in the future, if i do so, consider implementing sessions for this
		oauth2.AccessTypeOffline,
		oauth2.S256ChallengeOption(codeVerifier),
	)

	http.Redirect(w, r, url, http.StatusSeeOther)
}
