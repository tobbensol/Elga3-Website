package Review

import (
	"github.com/tobbensol/elga_3_website/internal/DB"
)

func (Review) Migrate() {
	if !DB.DB.Migrator().HasTable(&Review{}) {
		println("Table reviews not found, Creating table")
		err := DB.DB.AutoMigrate(&Review{})
		if err != nil {
			return
		}
	}
}

func (Review) GetReviews() []Review {
	var output []Review
	DB.DB.Find(&output)
	return output
}

func (Review) CreateReview(name string, score uint8) Review {
	review := Review{AlbumName: name, Score: score}
	DB.DB.Create(&review)
	return review
}

func (Review) DeleteAll() {
	DB.DB.Where("1=1").Delete(&Review{})
}
