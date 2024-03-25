package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Coordinate struct {
	X string
	Y string
}

func mainSite(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./html_templates/main_page.html"))

	coordinates := map[string][]Coordinate{
		"Coordinates": {
			{X: "1", Y: "2"},
			{X: "2", Y: "1"},
		},
	}

	tmpl.Execute(w, coordinates)
}

func add_coordinate(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./html_templates/main_page.html"))
	x := r.PostFormValue("X")
	y := r.PostFormValue("Y")
	tmpl.ExecuteTemplate(w, "coordinate", Coordinate{X: x, Y: y})
}

func get_db_connection_str() string{
	envFile, _ := godotenv.Read("../.env")
	return envFile["connection_str"]
}

func setup_db() (*gorm.DB, error){
	dsn := get_db_connection_str()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}

func main() {
	
	http.HandleFunc("/", mainSite)
	http.HandleFunc("/add_coordinate/", add_coordinate)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Server hosted at: http://localhost:8000/")

	db, _ := setup_db()
	fmt.Println(db.Name())

	log.Fatal(http.ListenAndServe(":8000", nil))
}
