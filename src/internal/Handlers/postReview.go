package Handlers

import (
	"fmt"
	"github.com/tobbensol/elga_3_website/UI/Templates"
	"github.com/tobbensol/elga_3_website/internal/Handlers/HandlerStructs"
	"github.com/tobbensol/elga_3_website/internal/Models/Item"
	"github.com/tobbensol/elga_3_website/internal/Models/Review"
	"github.com/tobbensol/elga_3_website/internal/Models/User"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

func PostReview(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		itemName := r.PostFormValue("ItemName")
		s := r.PostFormValue("Score")

		fmt.Println("test")

		errMessage := ""
		reviewData := HandlerStructs.ReviewData{}

		score, err := strconv.ParseUint(s, 10, 8)
		if err != nil {
			errMessage = "Score must be an integer"
		}

		user, err := User.GetUserFromCookie(db, r)
		if err != nil {
			errMessage = "Not logged in"
		}
		item, err := Item.FindByName(db, itemName)
		if err != nil {
			errMessage = "Item not found"
		}

		fmt.Println(errMessage)
		if errMessage == "" {
			review := Review.Create(db, user.ID, item.ID, uint8(score))
			reviewData = HandlerStructs.GetReviewData(db, review)
		}

		err = Templates.PostReview(reviewData, errMessage).Render(r.Context(), w)
		if err != nil {
			log.Printf("Error rendering template: %v", err)
		}
	}
}
