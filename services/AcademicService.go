package services

import (
	"github.com/zxcas321/ProfileGolang/mappers"
	"github.com/zxcas321/ProfileGolang/models"
	"github.com/zxcas321/ProfileGolang/repositories"
	"github.com/zxcas321/ProfileGolang/response"
)

func CreateAcademic(academic *models.Academics) error {
	return repositories.AcademicRepository.CreateAcademic(academic)
}

func FindAllAcademic() ([]response.AcademicsResponse, error) {
	academic, err := repositories.AcademicRepository.FindAllAcademic()
	if err != nil {
		return nil, err
	}
	res := make([]response.AcademicsResponse, 0, len(academic))
	for i := range academic{
		res = append(res, mappers.MapAcademicToResponse(academic[i]))
	}

	return res, nil
}

func FindByIDAcademic(id uint) (response.AcademicsResponse, error) {
	academic, err := repositories.AcademicRepository.FindByIDAcademic(id)
	if err != nil {
		return response.AcademicsResponse{}, err
	}

	return mappers.MapAcademicToResponse(academic), nil
}

func UpdateAcademic(id uint, data map[string]interface{}) error {
	return repositories.AcademicRepository.UpdateAcademic(id, data)
}

func DeleteAcademic(id uint) error {
	return repositories.AcademicRepository.DeleteAcademic(id)
}