package Handlers

import (
	"context"
	"github.com/tobbensol/elga_3_website/UI/Templates"
	"github.com/tobbensol/elga_3_website/internal/Models/Review"
	"github.com/tobbensol/elga_3_website/internal/Models/User"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func GetMainPage(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := &User.User{}
		loggedIn := false

		// Attempt to retrieve the "code_verifier" cookie
		userID, err := r.Cookie("authorization")
		if err != nil {
			// If cookie not found, set loggedIn to false
			log.Printf("Error retrieving cookie: %v", err)
		} else {
			// Cookie found, find the user
			user, err = User.FindByDiscordID(db, userID.Value)
			if err != nil {
				// if user doesn't exist, the user is not logged in
				log.Printf("Error retrieving user by Discord ID: %v", err)
			} else {
				loggedIn = true
			}
		}

		// Render the main page template
		reviews := Review.GetAll(db)
		if loggedIn {
			err = Templates.MainPage(reviews, loggedIn, *user).Render(context.Background(), w)
			if err != nil {
				log.Printf("Error rendering template: %v", err)
			}
		}
	}
}
