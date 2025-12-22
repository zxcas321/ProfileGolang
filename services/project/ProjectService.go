package project

import (
	"github.com/zxcas321/ProfileGolang/mappers"
	"github.com/zxcas321/ProfileGolang/models"
	"github.com/zxcas321/ProfileGolang/repositories"
	"github.com/zxcas321/ProfileGolang/response"
)

func CreateProject(project *models.Project) error {
	return repositories.ProjectRepository.CreateProject(project)
}

func FindAllProject() ([]response.ProjectResponse, error){
	projects, err := repositories.ProjectRepository.FindAllProject()
	if err != nil {
		return nil, err
	}
	res := make([]response.ProjectResponse, 0, len(projects))
	for i := range projects{
		res = append(res, mappers.MapProjectToResponse(projects[i]))
	}
	return res, nil
}

func FindByIDProject(id uint) (response.ProjectResponse, error){
	project, err := repositories.ProjectRepository.FindByIDProject(id)
	if err != nil {
		return response.ProjectResponse{}, err
	}
	
	return mappers.MapProjectToResponse(project), nil
} 

func UpdateProject(id uint, data map[string]interface{}) error{
	return repositories.ProjectRepository.UpdateProject(id, data)
}

func DeleteProject(id uint) error {
	return repositories.ProjectRepository.DeleteProject(id)
}