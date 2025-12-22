package admin

import (
	"github.com/zxcas321/ProfileGolang/mappers"
	"github.com/zxcas321/ProfileGolang/models"
	"github.com/zxcas321/ProfileGolang/repositories"
	"github.com/zxcas321/ProfileGolang/response"
)

func CreateAdmin(admin *models.Admin) error {
	return repositories.AdminRepository.CreateAdmin(admin)
}

func FindAllAdmin()([]response.AdminResponse,error) {
	admin, err := repositories.AdminRepository.FindAllAdmin()
	if err != nil {
		return nil, err
	}
	res := make([]response.AdminResponse, 0, len(admin))
	for i := range admin {
		res = append(res, mappers.MapAdminToResponse(admin[i]))
	}

	return res, nil
}

func FindByIDAdmin(id uint) (response.AdminResponse, error) {
	admin, err := repositories.AdminRepository.FindByIDAdmin(id)
	if err != nil {
		return response.AdminResponse{},err
	}

	return mappers.MapAdminToResponse(admin), nil
}

func UpdateAdmin(id uint, data map[string]interface{}) error {
	return repositories.AdminRepository.UpdateAdmin(id, data)
}

func DeleteAdmin(id uint) error {
	return repositories.AdminRepository.DeleteAdmin(id)
}