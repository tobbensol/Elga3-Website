package Server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/tobbensol/elga_3_website/internal/Auth"
	"github.com/tobbensol/elga_3_website/internal/Handlers"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func Connect(db *gorm.DB) {
	// mux will let me do neat routes in the future, right now it's equivalent to http.HandleFunc
	r := mux.NewRouter()
	// all the http request handlers
	r.HandleFunc("/", Handlers.GetMainPage(db))
	r.HandleFunc("/postReview/", Handlers.PostReview(db))

	r.HandleFunc("/auth/login", Auth.Login)
	r.HandleFunc("/auth/callback", Auth.Callback(db))

	// host static files (css files, eventually images?)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("UI/static"))))

	fmt.Println("Server hosted at: http://localhost:8000/")
	log.Fatal(http.ListenAndServe(":8000", r))
}
