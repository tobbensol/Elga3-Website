package DB

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/tobbensol/elga_3_website/internal/Models/Review"
	"github.com/tobbensol/elga_3_website/internal/Models/User"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getDBConnectionStr() string {
	envFile, _ := godotenv.Read("./.env")
	return envFile["connection_str"]
}

func SetupDB() *gorm.DB {
	var db *gorm.DB
	connectionStr := getDBConnectionStr()
	db, _ = gorm.Open(postgres.Open(connectionStr), &gorm.Config{})
	err := db.AutoMigrate(&User.User{}, &Review.Review{})
	if err != nil {
		fmt.Println("failed to migrate DB schemas")
	}
	return db
}
