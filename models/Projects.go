package models

import(
	"gorm.io/gorm"
)

type Project struct{
	gorm.Model
	UserID uint `json:"user_id"`
	Name string `json:"name"`
	Description string `json:"description"`
	WebUrl string `json:"web_url"`
	
	Skills []Skills `gorm:"many2many:project_skills"`
}