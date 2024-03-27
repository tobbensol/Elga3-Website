package Handlers

import (
	"html/template"
	"net/http"
)

type coordinate struct {
	X string
	Y string
}

func AddCoordinate(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./UI/templates/main_page.html"))
	x := r.PostFormValue("X")
	y := r.PostFormValue("Y")
	err := tmpl.ExecuteTemplate(w, "coordinate", coordinate{X: x, Y: y})
	if err != nil {
		return
	}
}

func MainSite(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./UI/templates/main_page.html"))

	coordinates := map[string][]coordinate{
		"Coordinates": {
			{X: "1", Y: "2"},
			{X: "2", Y: "1"},
		},
	}

	err := tmpl.Execute(w, coordinates)
	if err != nil {
		return
	}
}
