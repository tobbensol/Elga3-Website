package Handlers

import (
	"github.com/tobbensol/elga_3_website/UI/Templates"
	"github.com/tobbensol/elga_3_website/internal/Models/Item"
	"github.com/tobbensol/elga_3_website/internal/Models/User"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func GetMainPage(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		loggedIn := true
		user, err := User.GetUserFromCookie(db, r)

		if err != nil {
			loggedIn = false
		}

		itemData := Item.GetAll(db)

		// Render the main page template
		err = Templates.MainPage(itemData, loggedIn, user).Render(r.Context(), w)
		if err != nil {
			log.Printf("Error rendering template: %v", err)
		}
	}
}
