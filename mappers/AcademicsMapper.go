package mappers

import (
	"time"

	"github.com/zxcas321/ProfileGolang/models"
	"github.com/zxcas321/ProfileGolang/response"
	"github.com/zxcas321/ProfileGolang/utils"
)

func MapAcademicToResponse(academics models.Academics) response.AcademicsResponse {
	return  response.AcademicsResponse{
		ID: utils.EncodeID(academics.ID),
		Institution: academics.Institution,
		Degree: academics.Degree,
		Field: academics.Field,
		StartDate: academics.StartDate.Format(time.RFC3339),
		EndDate: academics.EndDate.Format(time.RFC3339),
	}
}