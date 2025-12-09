package repositories

import (
	"github.com/zxcas321/ProfileGolang/config"
	"github.com/zxcas321/ProfileGolang/models"
	
	"gorm.io/gorm"
)

type UserRepositoryStruct struct{}

var UserRepository = &UserRepositoryStruct{}

func (r *UserRepositoryStruct) CreateUser(user *models.User) error {
	if err := config.DB.Create(user).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepositoryStruct) FindAllUser() ([]models.User, error) {
	var user []models.User

	err := config.DB.
		Preload("Experience").
		Preload("Project.Skills").
		Preload("Academics").
		Preload("Certification").Find(&user).
		Error

	return user, err
}

func (r *UserRepositoryStruct) FindByIDUser(id uint) (models.User, error) {
	var user models.User

	err := config.DB.Preload("Experience").
		Preload("Project.Skills").
		Preload("Academics").
		Preload("Certification").
		First(&user, id).
		Error

	return user, err
}

func (r *UserRepositoryStruct) UpdateUser(id uint, data map[string]interface{}) error {
	err := config.DB.
		Model(&models.User{}).
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

func (r *UserRepositoryStruct) DeleteUser(id uint) error {
	return config.DB.Delete(&models.User{}, id).Error
}
