package App

import (
	"github.com/tobbensol/elga_3_website/internal/DB"
	"github.com/tobbensol/elga_3_website/internal/Server"
)

func StartApp() {
	db := DB.SetupDB()

	//Review.DeleteAll(db)

	Server.Connect(db)
}
