package models

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Name string `json:"name"`
	Email string `json:"email"`
	LinkedinUrl string`json:"linkedin_url"`
	GithubUrl string `json:"github_url"`
	Location string `json:"location"`
	Bio string `json:"bio"`
	ProfilePicture string `json:"profile_picture"`

	Experience []Experience `gorm:"foreignKey:UserID" json:"experience"`
	Project []Project `gorm:"foreignKey:UserID" json:"Project"`
}
