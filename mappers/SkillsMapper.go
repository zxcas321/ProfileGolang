package mappers

import (
	"github.com/zxcas321/ProfileGolang/models"
	"github.com/zxcas321/ProfileGolang/response"
	"github.com/zxcas321/ProfileGolang/utils"
)

func MapSkillsToResponse(skills models.Skills) response.SkillsResponse {
	return response.SkillsResponse{
		ID: utils.EncodeID(skills.ID),
		SkillName: skills.SkillName,
		Category: skills.Category,
	}
}