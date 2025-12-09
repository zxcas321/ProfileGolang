package repositories

import(
	"github.com/zxcas321/ProfileGolang/config"
	"github.com/zxcas321/ProfileGolang/models"
)

type AcademicRepositoryStruct struct{}
var AcademicRepository = &AcademicRepositoryStruct{}

func (r *AcademicRepositoryStruct) Create(academic *models.Academics) error {
	if err := config.DB.Create(academic).Error; err != nil {
		return err
	}
	return nil
}

func (r *AcademicRepositoryStruct) FindAll() ([]models.Academics, error){
	var academic []models.Academics
	if err := config.DB.Find(&academic).Error; err != nil {
		return nil, err
	}
	return academic, nil
}

func (r *AcademicRepositoryStruct) FindByID(academic *models.Academics, id uint) error{
	return config.DB.First(academic, id).Error
}

func (r *AcademicRepositoryStruct) Update(academic *models.Academics) error {
	if err := config.DB.Save(academic).Error; err != nil {
		return err
	}
	return nil
}

func (r *AcademicRepositoryStruct) Delete(id uint) error {
	return config.DB.Delete(&models.Academics{}, id).Error
}