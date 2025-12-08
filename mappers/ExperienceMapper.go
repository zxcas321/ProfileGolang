package mappers

import (
	"time"

	"github.com/zxcas321/ProfileGolang/models"
	"github.com/zxcas321/ProfileGolang/response"
	"github.com/zxcas321/ProfileGolang/utils"
)

func MapExperienceToResponse(experience models.Experience) response.ExperienceResponse{
	return response.ExperienceResponse{
		ID: utils.EncodeID(experience.ID),
		CompanyName: experience.CompanyName,
		Role: experience.Role,
		StartDate: experience.StartDate.Format(time.RFC3339),
		EndDate: experience.EndDate.Format(time.RFC3339),
	}
}