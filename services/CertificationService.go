package services

import (
	"github.com/zxcas321/ProfileGolang/mappers"
	"github.com/zxcas321/ProfileGolang/models"
	"github.com/zxcas321/ProfileGolang/repositories"
	"github.com/zxcas321/ProfileGolang/response"
)

func CreateCertification(certification *models.Certification) error {
	return repositories.CertificationRepository.CreateCertification(certification)
}

func FindAllCertification() ([]response.CertificationResponse, error) {
	certification, err := repositories.CertificationRepository.FindAllCertification()
	if err != nil {
		return nil, err
	}
	res := make([]response.CertificationResponse, 0, len(certification))
	for i := range certification{
		res = append(res, mappers.MapCertificationToResponse(certification[i]))
	}

	return res, nil
}

func FindByIDCertification(id uint) (response.CertificationResponse, error) {
	certification, err := repositories.CertificationRepository.FindByIDCertification(id)
	if err != nil {
		return response.CertificationResponse{}, err
	}

	return mappers.MapCertificationToResponse(certification), err
}

func UpdateCertification(id uint, data map[string]interface{}) error {
	return repositories.CertificationRepository.UpdateCertification(id, data)
}

func DeleteCertification(id uint) error {
	return repositories.CertificationRepository.DeleteCertification(id)
}