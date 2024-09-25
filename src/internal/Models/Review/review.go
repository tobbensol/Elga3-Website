package Review

import (
	"gorm.io/gorm"
)

func GetAll(db *gorm.DB) []Review {
	var output []Review
	db.Find(&output)
	return output
}

func Create(db *gorm.DB, userID uint, itemID uint, score uint8) Review {
	review := Review{
		ItemID: itemID,
		Score:  score,
		UserID: userID}
	db.Create(&review)
	return review
}

func DeleteAll(db *gorm.DB) {
	db.Where("1=1").Delete(&Review{})
}

func FindByID(db *gorm.DB, id string) (*Review, error) {
	var review Review
	result := db.Where("ID = ?", id).First(&review)
	if result.Error != nil {
		return &review, result.Error
	}
	db.Preload("Reviews").First(&review, review.ID)
	return &review, nil
}
