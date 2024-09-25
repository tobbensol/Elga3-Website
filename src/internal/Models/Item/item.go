package Item

import (
	"gorm.io/gorm"
)

func FindByID(db *gorm.DB, id uint) (*Item, error) {
	var item Item
	result := db.Where("ID = ?", id).First(&item)
	if result.Error != nil {
		return &item, result.Error
	}
	db.Preload("Reviews").First(&item, item.ID)
	return &item, nil
}

func FindByName(db *gorm.DB, name string) (*Item, error) {
	var item Item
	result := db.Where("Name = ?", name).First(&item)
	if result.Error != nil {
		return &item, result.Error
	}
	db.Preload("Reviews").First(&item, item.ID)
	return &item, nil
}
