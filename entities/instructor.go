package entities

import (
	"gorm.io/gorm"
)

type Instructor struct {
	gorm.Model

	FirstName   string
	Lastname    string
	Description string
}
