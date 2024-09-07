package Review

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	AlbumName string
	Score     uint8
	UserID    uint
}
