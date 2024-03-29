package Models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	AlbumName string
	Score     uint8
}

type User struct {
	gorm.Model
	Username string
	reviews  []Review
}
