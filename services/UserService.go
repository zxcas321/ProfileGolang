package services

import (
	"errors"

	"github.com/zxcas321/ProfileGolang/config"
	"github.com/zxcas321/ProfileGolang/mappers"
	"github.com/zxcas321/ProfileGolang/models"
	"github.com/zxcas321/ProfileGolang/response"
	"gorm.io/gorm"
)

func Create(user *models.User) error {
	return config.DB.Create(user).Error
}

func FindAll() ([]response.UserResponse, error) {
	var users []models.User

	err := config.DB.Preload("Experience").
		Preload("Project.Skills").
		Preload("Academics").
		Preload("Certification").Find(&users).Error

	if err != nil {
		return nil, err
	}

	res := make([]response.UserResponse, 0, len(users))
	for i := range users {
		res = append(res, mappers.MapUserToResponse(users[i]))
	}
	return res, nil
}

func FindByID(id uint) (response.UserResponse, error) {
	var user models.User

	err := config.DB.Preload("Experience").
		Preload("Project.Skills").
		Preload("Academics").
		Preload("Certification").
		First(&user, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.UserResponse{}, err
		}
		return response.UserResponse{}, err
	}

	return mappers.MapUserToResponse(user), nil
}

func Update(id uint, data map[string]interface{}) error {
	return config.DB.Model(&models.User{}).
		Where("id = ?", id).
		Updates(data).Error
}

func Delete(Id uint) error {
	return config.DB.Delete(&models.User{}, Id).Error
}
