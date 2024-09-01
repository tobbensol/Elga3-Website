package Handlers

import (
	"github.com/tobbensol/elga_3_website/internal/Models/Review"
	"html/template"
	"net/http"
	"strconv"
)

func PostReview(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./UI/templates/form.html"))
	name := r.PostFormValue("Album_Name")
	s := r.PostFormValue("Score")
	scoreIsInt := true

	score, err := strconv.ParseUint(s, 10, 8)
	if err != nil {
		scoreIsInt = false
	}

	tmpl.ExecuteTemplate(w, "score error", scoreIsInt)

	if scoreIsInt {
		review := Review.Review{}.CreateReview(name, uint8(score))
		err = tmpl.ExecuteTemplate(w, "Review", review)
		if err != nil {
			return
		}
	}
}
