package repositories

import(
	"github.com/zxcas321/ProfileGolang/config"
	"github.com/zxcas321/ProfileGolang/models"
)

type ExperienceRepositoryStruct struct{}
var ExperienceRepository = &ExperienceRepositoryStruct{}

func (r *ExperienceRepositoryStruct) Create(experience *models.Experience) error{
	if err := config.DB.Create(experience).Error; err != nil {
		return err
	}
	return nil
}

func (r *ExperienceRepositoryStruct) FindAll() ([]models.Experience, error) {
	var experience []models.Experience
	if err := config.DB.Find(experience).Error; err != nil {
		return nil, err
	}
	return experience, nil
}

func (r *ExperienceRepositoryStruct) FindByID(experience *models.Experience, id uint) error{
	return config.DB.First(experience, id).Error
}

func (r *ExperienceRepositoryStruct) Update(experience *models.Experience) error {
	if err := config.DB.Save(experience).Error; err != nil {
		return err
	}
	return nil
}

func (r *ExperienceRepositoryStruct) Delete(id uint) error {
	return config.DB.Delete(&models.Experience{}, id).Error
}