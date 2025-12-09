package repositories

import (
	"github.com/zxcas321/ProfileGolang/config"
	"github.com/zxcas321/ProfileGolang/models"
	"gorm.io/gorm"
)

type AdminRepositoryStruct struct{}

var AdminRepository = &AdminRepositoryStruct{}

func (r *AdminRepositoryStruct) CreateAdmin(admin *models.Admin) error {
	if err := config.DB.Create(admin).Error; err != nil {
		return err
	}
	return nil
}

func (r *AdminRepositoryStruct) FindAllAdmin() ([]models.Admin, error) {
	var admin []models.Admin

	err := config.DB.
		Find(&admin).
		Error

	return admin, err
}

func (r *AdminRepositoryStruct) FindByIDAdmin(id uint) (models.Admin, error) {
	var admin models.Admin

	err := config.DB.
		First(&admin, id).
		Error

	return admin, err
}

func (r *AdminRepositoryStruct) UpdateAdmin(id uint, data map[string]interface{}) error {
	err := config.DB.
		Model(&models.Admin{}).
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

func (R *AdminRepositoryStruct) DeleteAdmin(id uint) error {
	return config.DB.Delete(&models.Admin{}, id).Error
}
