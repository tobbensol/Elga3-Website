package Handlers

import (
	"fmt"
	"github.com/tobbensol/elga_3_website/internal/DB"
	"github.com/tobbensol/elga_3_website/internal/Models"
	"html/template"
	"net/http"
)

func MainSite(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./UI/templates/main_page.html"))

	var reviews []Models.Review

	DB.DB.Find(&reviews)

	fmt.Println(reviews)

	returnReviews := map[string][]Models.Review{
		"Reviews": reviews,
	}

	err := tmpl.Execute(w, returnReviews)
	if err != nil {
		return
	}
}
