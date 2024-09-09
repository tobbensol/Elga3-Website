package Auth

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"net/http"
)

func Callback(w http.ResponseWriter, r *http.Request) {
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

	http.SetCookie(w, &http.Cookie{
		Name:   "code_verifier",
		MaxAge: -1, // Delete the cookie
		Path:   "/",
	})

	// Exchange the authorization code for an access token
	token, err := oauthConfig.Exchange(context.Background(), code, oauth2.VerifierOption(verifier))
	if err != nil {
		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// You can now use the token to access the user's data
	// Here, we'll just display the token
	data := fmt.Sprintf("Access Token: %s\nRefresh Token: %s\n", token.AccessToken, token.RefreshToken)

	w.Write([]byte(data))
}
