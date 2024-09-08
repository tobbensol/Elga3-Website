package Auth

import (
	"context"
	"fmt"
	"net/http"
)

func Callback(w http.ResponseWriter, r *http.Request) {
	fmt.Println("this happens")
	state := r.URL.Query().Get("state")
	if state != "state" {
		fmt.Println("state is wrong")
		return
	}

	// Get the authorization code from the URL
	code := r.URL.Query().Get("code")

	// Exchange the authorization code for an access token
	token, err := getConfig().Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// You can now use the token to access the user's data
	// Here, we'll just display the token
	data := fmt.Sprintf("Access Token: %s\nRefresh Token: %s\n", token.AccessToken, token.RefreshToken)

	w.Write([]byte(data))
}
