package Handlers

import (
	"fmt"
	"github.com/tobbensol/elga_3_website/src/internal/Models/Review"
	"gorm.io/gorm"
	"html/template"
	"net/http"
	"strconv"
)

func PostReview(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./src/UI/templates/form.html"))
		name := r.PostFormValue("Album_Name")
		s := r.PostFormValue("Score")
		scoreIsInt := true

		score, err := strconv.ParseUint(s, 10, 8)
		if err != nil {
			scoreIsInt = false
		}

		err = tmpl.ExecuteTemplate(w, "score error", scoreIsInt)
		if err != nil {
			return
		}

		if scoreIsInt {
			review := Review.Create(db, name, uint8(score))
			err = tmpl.ExecuteTemplate(w, "review", review)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}
