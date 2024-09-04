package User

import (
	"gorm.io/gorm"
)

func Get(db *gorm.DB) []User {
	var output []User
	db.Find(&output)
	return output
}

func Create(db *gorm.DB, name string) User {
	User := User{Username: name}
	db.Create(&User)
	return User
}

func DeleteAll(db *gorm.DB) {
	db.Where("1=1").Delete(&User{})
}
