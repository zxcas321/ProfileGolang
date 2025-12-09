package repositories

import (
	"github.com/zxcas321/ProfileGolang/config"
	"github.com/zxcas321/ProfileGolang/models"

	"gorm.io/gorm"
)

type CertificationRepositoryStruct struct{}

var CertificationRepository = &CertificationRepositoryStruct{}

func (r *CertificationRepositoryStruct) CreateCertification(certification *models.Certification) error {
	if err := config.DB.Create(certification).Error; err != nil {
		return err
	}
	return nil
}

func (r *CertificationRepositoryStruct) FindAllCertification() ([]models.Certification, error) {
	var certification []models.Certification

	err := config.DB.
		Find(&certification).
		Error

	return certification, err
}

func (r *CertificationRepositoryStruct) FindByIDCertification(id uint) (models.Certification, error) {
	var certification models.Certification

	err := config.DB.
		First(&certification, id).
		Error

	return certification, err
}

func (r *CertificationRepositoryStruct) UpdateCertification(id uint, data map[string]interface{}) error {
	err := config.DB.
		Model(&models.Certification{}).
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

func (R *CertificationRepositoryStruct) DeleteCertification(id uint) error {
	return config.DB.Delete([]models.Academics{}, id).Error
}
