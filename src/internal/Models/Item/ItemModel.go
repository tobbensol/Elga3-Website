package Item

import (
	"github.com/tobbensol/elga_3_website/internal/Models/Review"
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name     string
	Category string
	IconPath string
	reviews  []Review.Review `gorm:"foreignkey:ItemID"`
}
