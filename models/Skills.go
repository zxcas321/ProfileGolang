package models

import (
	"gorm.io/gorm"
)

type Skills struct {
	gorm.Model
	ProjectID uint `json:"project_id"`
	SkillName string `json:"skill_name"`
	Category  string `json:"category"`
}
