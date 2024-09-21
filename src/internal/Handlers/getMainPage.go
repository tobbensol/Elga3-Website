package Handlers

import (
	"github.com/tobbensol/elga_3_website/src/internal/Models/Review"
	"gorm.io/gorm"
	"html/template"
	"net/http"
)

func GetMainPage(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./src/UI/templates/main_page.html"))

		returnReviews := map[string][]Review.Review{
			"Reviews": Review.GetAll(db), //db goes here
		}

		err := tmpl.Execute(w, returnReviews)
		if err != nil {
			return
		}
	}
}
