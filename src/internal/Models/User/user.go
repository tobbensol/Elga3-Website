package User

import (
	"github.com/tobbensol/elga_3_website/internal/Models/Review"
	"gorm.io/gorm"
)

func Get(db *gorm.DB) []User {
	var output []User
	db.Find(&output)
	return output
}

func Create(db *gorm.DB, name string) User {
	User := User{Username: name, Review: []Review.Review{}}
	db.Create(&User)
	return User
}

func DeleteAll(db *gorm.DB) {
	db.Where("1=1").Delete(&User{})
}
