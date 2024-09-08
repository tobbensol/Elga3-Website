package Server

import (
	"fmt"
	"github.com/tobbensol/elga_3_website/internal/Auth"
	"github.com/tobbensol/elga_3_website/internal/Handlers"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func Connect(db *gorm.DB) {
	// all the http request handlers
	http.HandleFunc("/", Handlers.GetMainPage(db))
	http.HandleFunc("/postReview/", Handlers.PostReview(db))

	http.HandleFunc("/auth/login", Auth.Login)
	http.HandleFunc("/auth/callback", Auth.Callback)

	// host static files (css files, eventually images?)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("UI/static"))))

	fmt.Println("Server hosted at: http://localhost:8000/")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
