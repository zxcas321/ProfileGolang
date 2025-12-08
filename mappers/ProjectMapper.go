package mappers

import (
	"github.com/zxcas321/ProfileGolang/models"
	"github.com/zxcas321/ProfileGolang/response"
	"github.com/zxcas321/ProfileGolang/utils"
)

func MapProjectToResponse(project models.Project) response.ProjectResponse{
	var skills []response.SkillsResponse
	for _, p := range project.Skills {
		skills = append(skills, MapSkillsToResponse(p))
	}
	return response.ProjectResponse{
		ID: utils.EncodeID(project.ID),
		Name: project.Name,
		Description: project.Description,
		WebUrl: project.WebUrl,
		Skills: skills,
	}
}