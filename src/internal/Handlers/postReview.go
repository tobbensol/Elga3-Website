package Handlers

import (
	"context"
	"github.com/tobbensol/elga_3_website/UI/Templates"
	"github.com/tobbensol/elga_3_website/internal/Models/Review"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

func PostReview(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.PostFormValue("Album_Name")
		s := r.PostFormValue("Score")
		scoreIsInt := true
		review := Review.Review{}

		score, err := strconv.ParseUint(s, 10, 8)
		if err != nil {
			scoreIsInt = false
		} else {
			review = Review.Create(db, name, uint8(score))
		}

		err = Templates.PostReview(review, scoreIsInt).Render(context.Background(), w)
		if err != nil {
			log.Printf("Error rendering template: %v", err)
		}

	}
}
