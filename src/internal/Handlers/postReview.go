package Handlers

import (
	"fmt"
	"github.com/tobbensol/elga_3_website/internal/Models"
	"github.com/tobbensol/elga_3_website/internal/Models/methods"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func AddCoordinate(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./UI/templates/main_page.html"))
	nam := r.PostFormValue("Album_Name")
	sco := r.PostFormValue("Score")

	Score, err := strconv.ParseFloat(sco, 64)
	if err != nil {
		log.Fatal(err)
		return
	}

	dbEntry := Models.Review{AlbumName: nam, Score: uint8(Score)}
	methods.CreateReview(dbEntry)

	fmt.Println()

	err = tmpl.ExecuteTemplate(w, "review", dbEntry)
	if err != nil {
		return
	}
}
