package Server

import (
	"fmt"
	"github.com/tobbensol/elga_3_website/internal/Handlers"
	"log"
	"net/http"
)

func init() {
	// all the http request handlers
	http.HandleFunc("/", Handlers.GetMainPage)
	http.HandleFunc("/postReview/", Handlers.PostReview)

	// host static files (css files, eventually images?)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("UI/static"))))
}

func Connect() {
	fmt.Println("Server hosted at: http://localhost:8000/")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
