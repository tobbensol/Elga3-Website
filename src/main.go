package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Coordinate struct {
	X string
	Y string
}

func mainSite(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./src/html_templates/main_page.html"))

	coordinates := map[string][]Coordinate{
		"Coordinates": {
			{X: "1", Y: "2"},
			{X: "2", Y: "1"},
		},
	}

	tmpl.Execute(w, coordinates)
}

func add_coordinate(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./src/html_templates/main_page.html"))
	x := r.PostFormValue("X")
	y := r.PostFormValue("Y")
	tmpl.ExecuteTemplate(w, "coordinate", Coordinate{X: x, Y: y})
}

func main() {
	fmt.Println("Hello World!")

	http.HandleFunc("/", mainSite)
	http.HandleFunc("/add_coordinate/", add_coordinate)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Fatal(http.ListenAndServe(":8000", nil))
}
