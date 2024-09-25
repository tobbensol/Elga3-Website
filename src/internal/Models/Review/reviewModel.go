package Review

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	Score  uint8
	ItemID uint
	UserID uint
}
