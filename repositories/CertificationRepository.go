package repositories

import (
	"github.com/zxcas321/ProfileGolang/config"
	"github.com/zxcas321/ProfileGolang/models"
)

type CertificationRepositoryStruct struct{}

var CertificationRepository = &CertificationRepositoryStruct{}

func (r *CertificationRepositoryStruct) Create(certification *models.Certification) error {
	if err := config.DB.Create(certification).Error; err != nil {
		return err
	}
	return nil
}

func (r *CertificationRepositoryStruct) FindAll() ([]models.Certification, error) {
	var certification []models.Certification
	if err := config.DB.Find(certification).Error; err != nil {
		return nil, err
	}
	return certification, nil
}

func (r *CertificationRepositoryStruct) FindByID(certification *models.Certification, id uint) error {
	return config.DB.First(certification, id).Error
}

func (r *CertificationRepositoryStruct) Update(certification *models.Certification) error {
	if err := config.DB.Save(certification).Error; err != nil {
		return err
	}
	return nil
}

func (R *CertificationRepositoryStruct) Delete(id uint) error {
	return config.DB.Delete([]models.Academics{}, id).Error
}
