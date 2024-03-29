package Models

import (
	"github.com/tobbensol/elga_3_website/internal/DB"
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	AlbumName string
	Score     uint8
}

func (Review) GetReviews() []Review {
	var output []Review
	DB.DB.Find(&output)
	return output
}

func (Review) CreateReview(name string, score uint8) Review {
	review := Review{AlbumName: name, Score: score}
	DB.DB.Create(review)
	return review
}

func (Review) DeleteAll() {
	DB.DB.Where("1=1").Delete(&Review{})
}
