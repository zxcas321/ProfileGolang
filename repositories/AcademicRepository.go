package repositories

import (
	"github.com/zxcas321/ProfileGolang/config"
	"github.com/zxcas321/ProfileGolang/models"

	"gorm.io/gorm"
)

type AcademicRepositoryStruct struct{}

var AcademicRepository = &AcademicRepositoryStruct{}

func (r *AcademicRepositoryStruct) CreateAcademic(academic *models.Academics) error {
	if err := config.DB.Create(academic).Error; err != nil {
		return err
	}
	return nil
}

func (r *AcademicRepositoryStruct) FindAllAcademic() ([]models.Academics, error) {
	var academic []models.Academics

	err := config.DB.
		Find(&academic).
		Error

	return academic, err
}

func (r *AcademicRepositoryStruct) FindByIDAcademic(id uint) (models.Academics, error) {
	var academic models.Academics

	err := config.DB.
		First(&academic, id).
		Error

	return academic, err
}

func (r *AcademicRepositoryStruct) UpdateAcademic(id uint, data map[string]interface{}) error {
	err := config.DB.
		Model(&models.Academics{}).
		Where("id = ?", id).
		Updates(data)

	if err.Error != nil {
		return err.Error
	}

	if err.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *AcademicRepositoryStruct) DeleteAcademic(id uint) error {
	return config.DB.Delete(&models.Academics{}, id).Error
}
