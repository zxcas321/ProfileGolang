package models

import(
	"time"

	"gorm.io/gorm"
)

type Experience struct{
	gorm.Model
	UserID uint `json:"user_id"`
	CompanyName string `json:"company_name"`
	Role string `json:"role"`
	StartDate time.Time `json:"start_date"`
	EndDate time.Time `json:"end_date"`
}