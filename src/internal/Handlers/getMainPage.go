package Handlers

import (
	"context"
	"fmt"
	"github.com/tobbensol/elga_3_website/src/UI/Templates"
	"github.com/tobbensol/elga_3_website/src/internal/Models/Review"
	"github.com/tobbensol/elga_3_website/src/internal/Models/User"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func GetMainPage(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var loggedIn bool
		var user *User.User // Assuming your User model is named User and this is a pointer.

		// Attempt to retrieve the "code_verifier" cookie
		userID, err := r.Cookie("authorization")
		if err != nil {
			// If there's an error (e.g., cookie not found), set loggedIn to false and log the error
			loggedIn = false
			log.Printf("Error retrieving cookie: %v", err)
		} else {
			// Cookie found, now attempt to find the user by their Discord ID
			user, err = User.FindByDiscordID(db, userID.Value)
			if err != nil {
				// Handle the case where user is not found or there's an error
				loggedIn = false
				log.Printf("Error retrieving user by Discord ID: %v", err)
			} else {
				loggedIn = true
			}
		}
		fmt.Println(user.GetAvatar())

		// Render the main page template, passing the reviews and user status
		reviews := Review.GetAll(db)
		if loggedIn {
			err = Templates.Main(reviews, loggedIn, *user).Render(context.Background(), w)
		} else {
			err = Templates.Main(reviews, loggedIn, User.User{}).Render(context.Background(), w)
		}
	}
}
