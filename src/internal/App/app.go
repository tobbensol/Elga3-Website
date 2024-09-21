package App

import (
	"github.com/tobbensol/elga_3_website/src/internal/DB"
	"github.com/tobbensol/elga_3_website/src/internal/Models/Review"
	"github.com/tobbensol/elga_3_website/src/internal/Server"
)

func StartApp() {
	db := DB.SetupDB()

	Review.DeleteAll(db)

	Server.Connect(db)
}
