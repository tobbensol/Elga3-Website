package User

import (
	"github.com/tobbensol/elga_3_website/internal/Models/Review"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username      string
	DiscordID     string
	DiscordAvatar string
	AccessToken   string
	RefreshToken  string
	TokenExpiry   time.Time
	Reviews       []Review.Review
}
