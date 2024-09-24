package Review

import (
	"gorm.io/gorm"
)

func GetAll(db *gorm.DB) []Review {
	var output []Review
	db.Find(&output)
	return output
}

func Create(db *gorm.DB, userID uint, name string, score uint8) Review {
	review := Review{
		AlbumName: name,
		Score:     score,
		UserID:    userID}
	db.Create(&review)
	return review
}

func DeleteAll(db *gorm.DB) {
	db.Where("1=1").Delete(&Review{})
}
