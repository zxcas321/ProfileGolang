package models

import (
	"gorm.io/gorm"
)

type Skills struct{
	gorm.Model
	SkillName string `json:"skill_name"`
	Category string `json:"category"`
}