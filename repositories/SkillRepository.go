package repositories

import(
	"github.com/zxcas321/ProfileGolang/config"
	"github.com/zxcas321/ProfileGolang/models"
)

type SkillRepositoryStruct struct{}
var SkillRepository = &SkillRepositoryStruct{}

func (r *SkillRepositoryStruct) Create(skill *models.Skills) error{
	if err := config.DB.Create(skill).Error; err != nil{
		return err
	}
	return nil
}

func (r *SkillRepositoryStruct) FindAll() ([]models.Skills, error){
	var skill []models.Skills
	if err := config.DB.Find(&skill).Error; err != nil{
		return nil, err
	}
	return skill, nil
}

func (r *SkillRepositoryStruct) FindByID(skill *models.Skills, id uint) error{
	return config.DB.First(skill, id).Error
}

func (r *SkillRepositoryStruct) Update(skill *models.Skills) error{
	if err := config.DB.Save(skill).Error; err != nil{
		return err
	}
	return nil
}

func (r *SkillRepositoryStruct) Delete(id uint) error{
	return config.DB.Delete(&models.Skills{}, id).Error
}