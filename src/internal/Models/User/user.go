package User

import (
	"github.com/tobbensol/elga_3_website/internal/DB"
)

func (User) Migrate() {
	if !DB.DB.Migrator().HasTable(&User{}) {
		println("Table reviews not found, Creating table")
		err := DB.DB.AutoMigrate(&User{})
		if err != nil {
			return
		}
	}
}

func (User) GetUsers() []User {
	var output []User
	DB.DB.Find(&output)
	return output
}

func (User) Create(name string) User {
	User := User{Username: name}
	DB.DB.Create(&User)
	return User
}

func (User) DeleteAll() {
	DB.DB.Where("1=1").Delete(&User{})
}
