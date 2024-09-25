package HandlerStructs

import (
	"github.com/tobbensol/elga_3_website/internal/Models/Item"
	"github.com/tobbensol/elga_3_website/internal/Models/Review"
	"github.com/tobbensol/elga_3_website/internal/Models/User"
	"gorm.io/gorm"
)

type ReviewData struct {
	Review Review.Review
	User   *User.User
	Item   *Item.Item
}

func GetReviewData(db *gorm.DB, review Review.Review) ReviewData {
	item, _ := Item.FindByID(db, review.ItemID)
	user, _ := User.FindByID(db, review.UserID)

	return ReviewData{
		Review: review,
		User:   user,
		Item:   item,
	}
}

func GetAllReviewData(db *gorm.DB, reviews []Review.Review) []ReviewData {
	var reviewDetailsList []ReviewData
	for _, review := range reviews {
		reviewDetailsList = append(reviewDetailsList, GetReviewData(db, review))
	}
	return reviewDetailsList
}
