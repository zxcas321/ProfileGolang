package repositories

import (
	"github.com/zxcas321/ProfileGolang/config"
	"github.com/zxcas321/ProfileGolang/models"
	"gorm.io/gorm"
)

type SkillRepositoryStruct struct{}

var SkillRepository = &SkillRepositoryStruct{}

func (r *SkillRepositoryStruct) CreateSkill(skill *models.Skills) error {
	if err := config.DB.Create(skill).Error; err != nil {
		return err
	}
	
	return nil
}

func (r *SkillRepositoryStruct) FindAllSkill() ([]models.Skills, error) {
	var skill []models.Skills

	err := config.DB.Find(&skill).Error

	return skill, err
}

func (r *SkillRepositoryStruct) FindByIDSkill(id uint) (models.Skills, error) {
	var skill models.Skills
	err := config.DB.
		First(&skill, id).
		Error

	return skill, err
}

func (r *SkillRepositoryStruct) UpdateSkill(id uint, data map[string]interface{}) error {
	err := config.DB.
		Model(&models.Skills{}).
		Where("id = ?", id).
		Updates(data)
	
	if err.Error != nil{
		return err.Error
	}

	if err.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *SkillRepositoryStruct) DeleteSkill(id uint) error {
	return config.DB.Delete(&models.Skills{}, id).Error
}
