package models

import (
	"time"

	"gorm.io/gorm"
)

type Certification struct {
	gorm.Model
	UserID     uint       `json:"user_id"`
	Title      string     `json:"title"`
	Issuer     string     `json:"issuer"`
	IssueDate  time.Time  `json:"issue_date"`
	ExpiryDate *time.Time `json:"expiry_date,omitempty"`
	ImageUrl   string     `json:"image_url"`
}
