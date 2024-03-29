package methods

import (
	"github.com/tobbensol/elga_3_website/internal/DB"
	"github.com/tobbensol/elga_3_website/internal/Models"
)

func GetReviews() []Models.Review {
	var output []Models.Review
	DB.DB.Find(&output)
	return output
}

func CreateReview(review Models.Review) {
	DB.DB.Create(&review)
}

func DeleteAll() {
	DB.DB.Where("1=1").Delete(&Models.Review{})
}
