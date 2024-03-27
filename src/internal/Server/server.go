package Server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tobbensol/elga_3_website/internal/DB"
	"github.com/tobbensol/elga_3_website/internal/Handlers"
)

func InitServer() {
	http.HandleFunc("/", Handlers.MainSite)
	http.HandleFunc("/add_coordinate/", Handlers.AddCoordinate)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("UI/static"))))
	fmt.Println("Server hosted at: http://localhost:8000/")

	db, _ := DB.SetupDb()
	fmt.Println(db.Name())

	log.Fatal(http.ListenAndServe(":8000", nil))
}
