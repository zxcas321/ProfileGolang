package repositories

import(
	"github.com/zxcas321/ProfileGolang/config"
	"github.com/zxcas321/ProfileGolang/models"
)

type ProjectRepositoryStruct struct{}
var ProjectRepository = &ProjectRepositoryStruct{}

func (r *ProjectRepositoryStruct) Create(project *models.Project) error{
	if err := config.DB.Create(&project).Error; err != nil {
		return err
	}
	return nil
}

func (r *ProjectRepositoryStruct) FindAll()([]models.Project, error) {
	var project []models.Project
	if err := config.DB.Find(&project).Error; err != nil{
		return nil,err
	}
	return project, nil
}

func (r *ProjectRepositoryStruct) FindByID(project *models.Project, id uint) error {
	return config.DB.First(project, id).Error
}

func (r *ProjectRepositoryStruct) Update(project *models.Project) error {
	if err := config.DB.Save(project).Error; err != nil {
		return err
	}
	return nil
}

func (r *ProjectRepositoryStruct) Delete(id uint) error {
	return config.DB.Delete(&models.Project{}, id).Error
}
