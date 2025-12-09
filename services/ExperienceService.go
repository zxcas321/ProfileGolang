package services

import (
	"github.com/zxcas321/ProfileGolang/mappers"
	"github.com/zxcas321/ProfileGolang/models"
	"github.com/zxcas321/ProfileGolang/repositories"
	"github.com/zxcas321/ProfileGolang/response"
)

func CreateExperience(experience *models.Experience) error {
	return repositories.ExperienceRepository.CreateExperience(experience)
}

func FindAllExperience() ([]response.ExperienceResponse, error){
	experience, err := repositories.ExperienceRepository.FindAllExperience()
	if err != nil {
		return nil, err
	}
	res := make([]response.ExperienceResponse, 0, len(experience))
	for i :=range experience{
		res = append(res, mappers.MapExperienceToResponse(experience[i]))
	}

	return res, nil
}

func FindByIDExperience(id uint) (response.ExperienceResponse, error){
	experience, err := repositories.ExperienceRepository.FindByIDExperience(id)
	if err != nil {
		return response.ExperienceResponse{}, err
	}

	return mappers.MapExperienceToResponse(experience), nil
}

func UpdateExperience(id uint, data map[string]interface{}) error {
	return repositories.ExperienceRepository.UpdateExperience(id, data)
}

func DeleteExperience(id uint) error {
	return repositories.ExperienceRepository.DeleteExperience(id)
}