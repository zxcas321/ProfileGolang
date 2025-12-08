package models

type ProjectSkills struct{
	ProjectID uint `gorm:"primaryKey"`
	SkillsID uint `gorm:"primaryKey"`
	Level string `json:"level"`
}