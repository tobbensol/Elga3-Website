package User

import (
	"github.com/tobbensol/elga_3_website/internal/Models/Review"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Review   []Review.Review
}
