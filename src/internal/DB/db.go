package DB

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getDbConnectionStr() string {
	envFile, _ := godotenv.Read("./.env")
	return envFile["connection_str"]
}

func SetupDb() (*gorm.DB, error) {
	dsn := getDbConnectionStr()
	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}
