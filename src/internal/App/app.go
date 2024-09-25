package App

import (
	"github.com/tobbensol/elga_3_website/internal/DB"
	"github.com/tobbensol/elga_3_website/internal/Models/Item"
	"github.com/tobbensol/elga_3_website/internal/Models/Review"
	"github.com/tobbensol/elga_3_website/internal/Models/User"
	"github.com/tobbensol/elga_3_website/internal/Server"
)

func StartApp() {
	db := DB.SetupDB()

	Review.DeleteAll(db)
	User.DeleteAll(db)
	Item.DeleteAll(db)

	Server.Connect(db)
}
