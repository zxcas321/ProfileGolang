package repositories

import (
	"github.com/zxcas321/ProfileGolang/config"
	"github.com/zxcas321/ProfileGolang/models"
	"gorm.io/gorm"
)

type ProjectRepositoryStruct struct{}

var ProjectRepository = &ProjectRepositoryStruct{}

func (r *ProjectRepositoryStruct) CreateProject(project *models.Project) error {
	if err := config.DB.Create(project).Error; err != nil {
		return err
	}

	return nil
}

func (r *ProjectRepositoryStruct) FindAllProject() ([]models.Project, error) {
	var project []models.Project

	err := config.DB.
		Preload("Skills").
		Find(&project).
		Error

	return project, err
}

func (r *ProjectRepositoryStruct) FindByIDProject(id uint) (models.Project, error) {
	var project models.Project

	err := config.DB.
		Preload("Skills").
		First(&project, id).
		Error

	return project, err
}

func (r *ProjectRepositoryStruct) UpdateProject(id uint, data map[string]interface{}) error {
	err := config.DB.
		Model(&models.Project{}).
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

func (r *ProjectRepositoryStruct) DeleteProject(id uint) error {
	return config.DB.Delete(&models.Project{}, id).Error
}
