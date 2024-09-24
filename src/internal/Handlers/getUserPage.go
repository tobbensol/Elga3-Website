package Handlers

import (
	"github.com/go-chi/chi"
	"github.com/tobbensol/elga_3_website/UI/Templates"
	"github.com/tobbensol/elga_3_website/internal/Models/User"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func GetUserPage(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		loggedIn := true
		user, err := User.GetUserFromCookie(db, r)
		if err != nil {
			loggedIn = false
		}
		userPage, err := User.FindByDiscordUsername(db, chi.URLParam(r, "user"))
		if err != nil {
			// throw 404 error
		}

		reviews := userPage.Reviews

		// Render the main page template
		err = Templates.MainPage(reviews, loggedIn, user).Render(r.Context(), w)
		if err != nil {
			log.Printf("Error rendering template: %v", err)
		}
	}
}
