package Handlers

import (
	"fmt"
	"github.com/tobbensol/elga_3_website/UI/Templates"
	"github.com/tobbensol/elga_3_website/internal/Models/Review"
	"github.com/tobbensol/elga_3_website/internal/Models/User"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

func PostReview(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.PostFormValue("Album_Name")
		s := r.PostFormValue("Score")

		fmt.Println("test")

		errMessage := ""
		review := Review.Review{}

		score, err := strconv.ParseUint(s, 10, 8)
		if err != nil {
			errMessage = "Score must be an integer"
		}

		user, err := User.GetUserFromCookie(db, r)
		if err != nil {
			errMessage = "Not logged in"
		}

		fmt.Println(errMessage)
		if errMessage == "" {
			review = Review.Create(db, user.ID, name, uint8(score))
		}

		err = Templates.PostReview(review, errMessage).Render(r.Context(), w)
		if err != nil {
			log.Printf("Error rendering template: %v", err)
		}

	}
}
