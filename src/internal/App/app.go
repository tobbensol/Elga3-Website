package App

import (
	"github.com/tobbensol/elga_3_website/internal/DB"
	"github.com/tobbensol/elga_3_website/internal/Models"
	"github.com/tobbensol/elga_3_website/internal/Server"
)

func StartApp() {
	DB.SetupDB()

	Models.Review{}.DeleteAll()

	Server.Connect()
}
