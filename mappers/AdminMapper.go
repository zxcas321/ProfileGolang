package mappers

import (
	"github.com/zxcas321/ProfileGolang/models"
	"github.com/zxcas321/ProfileGolang/response"
	"github.com/zxcas321/ProfileGolang/utils"
)

func MapAdminToResponse(admin models.Admin) response.AdminResponse{
	return response.AdminResponse{
		ID: utils.EncodeID(admin.ID),
		Email: admin.Email,
	}
}