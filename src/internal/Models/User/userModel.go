package User

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
}
