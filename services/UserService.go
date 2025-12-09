package services

import (
	"github.com/zxcas321/ProfileGolang/mappers"
	"github.com/zxcas321/ProfileGolang/models"
	"github.com/zxcas321/ProfileGolang/repositories"
	"github.com/zxcas321/ProfileGolang/response"
)

func CreateUser(user *models.User) error {
	return repositories.UserRepository.CreateUser(user)
}

func FindAllUser() ([]response.UserResponse, error) {
	users, err := repositories.UserRepository.FindAllUser()
	if err != nil {
		return nil, err
	}
	res := make([]response.UserResponse, 0, len(users))
	for i := range users {
		res = append(res, mappers.MapUserToResponse(users[i]))
	}

	return res, nil
}

func FindByIDUser(id uint) (response.UserResponse, error) {
	user, err := repositories.UserRepository.FindByIDUser(id)
	if err != nil {
		return response.UserResponse{}, err
	}

	return mappers.MapUserToResponse(user), nil
}

func UpdateUser(id uint, data map[string]interface{}) error {
	return repositories.UserRepository.UpdateUser(id, data)
}

func DeleteUser(id uint) error {
	return repositories.UserRepository.DeleteUser(id)
}
