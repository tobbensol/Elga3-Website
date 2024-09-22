package Server

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/tobbensol/elga_3_website/internal/Auth"
	"github.com/tobbensol/elga_3_website/internal/Handlers"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func Connect(db *gorm.DB) {
	// chi will let me do neat routes in the future, right now it's equivalent to http.HandleFunc
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	// all the http request handlers
	r.Get("/", Handlers.GetMainPage(db))
	r.Post("/postReview/", Handlers.PostReview(db))

	r.Get("/auth/login", Auth.Login)
	r.Get("/auth/callback", Auth.Callback(db))

	// host static files (css files, eventually images?)
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("src/UI/static"))))

	fmt.Println("Server hosted at: http://localhost:8000/")
	log.Fatal(http.ListenAndServe(":8000", r))
}
