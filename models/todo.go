package models

import (
	"gorm.io/gorm"
	"time"
)

type Todo struct {
	gorm.Model
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
