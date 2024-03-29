package DB

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func getDBConnectionStr() string {
	envFile, _ := godotenv.Read("./.env")
	return envFile["connection_str"]
}

func SetupDB() {
	dsn := getDBConnectionStr()
	DB, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})

}
