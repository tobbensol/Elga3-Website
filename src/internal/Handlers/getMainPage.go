package Handlers

import (
	"github.com/tobbensol/elga_3_website/UI/Templates"
	"github.com/tobbensol/elga_3_website/internal/Handlers/HandlerStructs"
	"github.com/tobbensol/elga_3_website/internal/Models/Review"
	"github.com/tobbensol/elga_3_website/internal/Models/User"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func GetMainPage(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		loggedIn := true
		reviews := Review.GetAll(db)
		user, err := User.GetUserFromCookie(db, r)

		if err != nil {
			loggedIn = false
		}

		reviewWithData := HandlerStructs.GetAllReviewData(db, reviews)

		// Render the main page template
		err = Templates.MainPage(reviewWithData, loggedIn, user).Render(r.Context(), w)
		if err != nil {
			log.Printf("Error rendering template: %v", err)
		}
	}
}
