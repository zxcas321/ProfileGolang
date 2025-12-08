package repositories

import(
	"github.com/zxcas321/ProfileGolang/config"
	"github.com/zxcas321/ProfileGolang/models"
)

type UserRepositoryStruct struct{}
var UserRepository = &UserRepositoryStruct{}

func (r *UserRepositoryStruct) Create(user *models.User) error{
	if err := config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryStruct) FindAll() ([]models.User, error){
	var user []models.User
	if err := config.DB.Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepositoryStruct) FindByID(user *models.User, id uint) error{
	return config.DB.First(user, id).Error
}

func (r *UserRepositoryStruct) Update(user *models.User) error{
	if err := config.DB.Save(user).Error; err != nil{
		return err
	}
	return nil
}

func (r *UserRepositoryStruct) Delete(id uint) error{
	return config.DB.Delete(&models.User{}, id).Error
}