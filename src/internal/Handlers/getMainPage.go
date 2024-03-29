package Handlers

import (
	"github.com/tobbensol/elga_3_website/internal/Models"
	"github.com/tobbensol/elga_3_website/internal/Models/methods"
	"html/template"
	"net/http"
)

func MainSite(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./UI/templates/main_page.html"))

	returnReviews := map[string][]Models.Review{
		"Reviews": methods.GetReviews(),
	}

	err := tmpl.Execute(w, returnReviews)
	if err != nil {
		return
	}
}
