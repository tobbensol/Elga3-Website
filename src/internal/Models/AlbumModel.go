package Models

import "gorm.io/gorm"

type Album struct {
	gorm.Model
	Name      string
	MeanScore float32
	ArtistID  uint
}
