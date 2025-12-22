package project

import (
	"github.com/zxcas321/ProfileGolang/mappers"
	"github.com/zxcas321/ProfileGolang/models"
	"github.com/zxcas321/ProfileGolang/repositories"
	"github.com/zxcas321/ProfileGolang/response"
)

func CreateSkill(skill *models.Skills) error {
	return repositories.SkillRepository.CreateSkill(skill)
}

func FindAllSkill() ([]response.SkillsResponse, error) {
	skills, err := repositories.SkillRepository.FindAllSkill()
	if err != nil {
		return nil, err
	}

	res := make([]response.SkillsResponse, 0, len(skills))
	for i := range skills {
		res = append(res, mappers.MapSkillsToResponse(skills[i]))
	}
	
	return res, nil
}

func FindByIDSkill(id uint) (response.SkillsResponse, error){
	skills, err := repositories.SkillRepository.FindByIDSkill(id)
	if err != nil {
		return response.SkillsResponse{}, err
	}

	return mappers.MapSkillsToResponse(skills), nil
}

func UpdateSkill(id uint, data map[string]interface{}) error {
	return repositories.SkillRepository.UpdateSkill(id, data)
}

func DeleteSkill(id uint) error {
	return repositories.SkillRepository.DeleteSkill(id)
}