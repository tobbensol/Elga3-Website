package App

import (
	"github.com/tobbensol/elga_3_website/internal/DB"
	"github.com/tobbensol/elga_3_website/internal/Models/methods"
	"github.com/tobbensol/elga_3_website/internal/Server"
)

func StartApp() {
	DB.SetupDB()

	methods.DeleteAll()

	Server.InitServer()
}
