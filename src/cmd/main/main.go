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
	tmpl := template.Must(template.ParseFiles("./templates/main_page.html"))

	coordinates := map[string][]Coordinate{
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

func addCoordinate(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/main_page.html"))
	x := r.PostFormValue("X")
	y := r.PostFormValue("Y")
	err := tmpl.ExecuteTemplate(w, "coordinate", Coordinate{X: x, Y: y})
	if err != nil {
		return
	}
}

func getDbConnectionStr() string {
	envFile, _ := godotenv.Read("../.env")
	return envFile["connection_str"]
}

func setupDb() (*gorm.DB, error) {
	dsn := getDbConnectionStr()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}

func main() {
	http.HandleFunc("/", mainSite)
	http.HandleFunc("/add_coordinate/", addCoordinate)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Server hosted at: http://localhost:8000/")

	db, _ := setupDb()
	fmt.Println(db.Name())

	log.Fatal(http.ListenAndServe(":8000", nil))
}
