package User

import (
	"github.com/tobbensol/elga_3_website/internal/Models/Review"
	"gorm.io/gorm"
	"time"
)

func Get(db *gorm.DB) []User {
	var output []User
	db.Find(&output)
	return output
}

func Create(db *gorm.DB, username string, discordID string, discordAvatar string, accessToken string, refreshToken string, tokenExpiry time.Time) User {
	User := User{
		Username:      username,
		DiscordID:     discordID,
		Review:        []Review.Review{},
		DiscordAvatar: discordAvatar,
		AccessToken:   accessToken,
		RefreshToken:  refreshToken,
		TokenExpiry:   tokenExpiry,
	}
	db.Create(&User)
	return User
}

func DeleteAll(db *gorm.DB) {
	db.Where("1=1").Delete(&User{})
}

func FindByDiscordID(db *gorm.DB, discordID string) (User, error) {
	var user User
	result := db.Where("discord_id = ?", discordID).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func DiscordUserExists(db *gorm.DB, discordID string) (bool, error) {
	var exists bool
	err := db.Model(&User{}).
		Select("count(*) > 0").
		Where("discord_id = ?", discordID).
		Find(&exists).
		Error
	if err != nil {
		return exists, err
	}
	return exists, nil
}
