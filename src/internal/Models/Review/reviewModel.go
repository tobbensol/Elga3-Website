package Review

import (
	"github.com/tobbensol/elga_3_website/internal/Models/User"
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	AlbumName string
	Score     uint8
	UserID    uint
	User      User.User `gorm:"foreignKey:UserID"`
}
