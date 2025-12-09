package repositories

import(
	"github.com/zxcas321/ProfileGolang/config"
	"github.com/zxcas321/ProfileGolang/models"

	"gorm.io/gorm"
)

type ExperienceRepositoryStruct struct{}
var ExperienceRepository = &ExperienceRepositoryStruct{}

func (r *ExperienceRepositoryStruct) CreateExperience(experience *models.Experience) error{
	if err := config.DB.Create(experience).Error; err != nil {
		return err
	}
	return nil
}

func (r *ExperienceRepositoryStruct) FindAllExperience() ([]models.Experience, error) {
	var experience []models.Experience
	err := config.DB.
	Find(&experience).
	Error

	return experience, err
}

func (r *ExperienceRepositoryStruct) FindByIDExperience(id uint) (models.Experience, error){
	var experience models.Experience
	
	err := config.DB.
		First(&experience, id).
		Error
	
	return experience, err
}

func (r *ExperienceRepositoryStruct) UpdateExperience(id uint, data map[string]interface{}) error {
	err := config.DB.
		Model(&models.Experience{}).
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

func (r *ExperienceRepositoryStruct) DeleteExperience(id uint) error {
	return config.DB.Delete(&models.Experience{}, id).Error
}