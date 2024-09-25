package User

import (
	"fmt"
	"github.com/tobbensol/elga_3_website/internal/Models/Review"
	"gorm.io/gorm"
	"log"
	"net/http"
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
		DiscordAvatar: discordAvatar,
		AccessToken:   accessToken,
		RefreshToken:  refreshToken,
		TokenExpiry:   tokenExpiry,
		Reviews:       []Review.Review{},
	}
	db.Create(&User)
	return User
}

func DeleteAll(db *gorm.DB) {
	db.Where("1=1").Delete(&User{})
}

func FindByDiscordID(db *gorm.DB, discordID string) (*User, error) {
	var user User
	result := db.Where("discord_id = ?", discordID).First(&user)
	if result.Error != nil {
		return &user, result.Error
	}
	db.Preload("Reviews").First(&user, user.ID)
	return &user, nil
}

func FindByID(db *gorm.DB, id uint) (*User, error) {
	var user User
	result := db.Where("ID = ?", id).First(&user)
	if result.Error != nil {
		return &user, result.Error
	}
	db.Preload("Reviews").First(&user, user.ID)
	return &user, nil
}

func DiscordUserExists(db *gorm.DB, discordID string) bool {
	var exists bool
	db.Model(&User{}).
		Select("count(*) > 0").
		Where("discord_id = ?", discordID).
		Find(&exists)
	return exists
}

func FindByDiscordUsername(db *gorm.DB, username string) (*User, error) {
	var user User
	result := db.Where("Username = ?", username).First(&user)
	if result.Error != nil {
		return &user, result.Error
	}
	db.Preload("Reviews").First(&user, user.ID)
	return &user, nil
}

func (user User) GetAvatar() string {
	if user.DiscordID == "" {
		return ""
	}
	avatar := fmt.Sprintf("https://cdn.discordapp.com/avatars/%s/%s.jpg", user.DiscordID, user.DiscordAvatar)
	return avatar
}

func GetUserFromCookie(db *gorm.DB, r *http.Request) (*User, error) {
	user := &User{}
	// Attempt to retrieve the "code_verifier" cookie
	userID, err := r.Cookie("authorization")
	if err == nil {
		// Cookie found, find the user
		user, err = FindByDiscordID(db, userID.Value)
		if err != nil {
			// if user doesn't exist, the user is not logged in
			log.Printf("Error retrieving user by Discord ID: %v", err)
		} else {
			return user, nil
		}
	}
	db.Preload("Reviews").First(&user, user.ID)
	return user, err
}
