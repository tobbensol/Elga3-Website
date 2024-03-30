package Handlers

import (
	"github.com/tobbensol/elga_3_website/internal/Models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func PostReview(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./UI/templates/main_page.html"))
	name := r.PostFormValue("Album_Name")
	s := r.PostFormValue("Score")

	score, err := strconv.ParseUint(s, 10, 8)
	if err != nil {
		log.Fatal(err)
		return
	}

	review := Models.Review{}.CreateReview(name, uint8(score))
	err = tmpl.ExecuteTemplate(w, "review", review)
	if err != nil {
		return
	}
}
