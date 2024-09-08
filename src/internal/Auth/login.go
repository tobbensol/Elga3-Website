package Auth

import (
	"golang.org/x/oauth2"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	oauth2Config := getConfig()
	// access type online = access even if the user is away
	url := oauth2Config.AuthCodeURL("state", oauth2.AccessTypeOffline)

	http.Redirect(w, r, url, http.StatusSeeOther)
}
