package models

import (
	"time"

	"gorm.io/gorm"
)

type Academics struct {
	gorm.Model
	UserID      uint      `json:"user_id"`
	Institution string    `json:"institution"`
	Degree      string    `json:"degree"`
	Field       string    `json:"field"` //computer science
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	GPA         string    `json:"GPA"`
	Description string    `json:"description"`
}
