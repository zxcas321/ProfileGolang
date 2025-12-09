package repositories

import(
	"github.com/zxcas321/ProfileGolang/config"
	"github.com/zxcas321/ProfileGolang/models"
)

type AdminRepositoryStruct struct{}
var AdminRepository = &AdminRepositoryStruct{}

func (r *AdminRepositoryStruct) Create(admin *models.Admin) error {
	if err := config.DB.Create(admin).Error; err != nil {
		return  err
	}
	return nil
}

func (r *AdminRepositoryStruct) FindAll() ([]models.Admin, error) {
	var admin []models.Admin
	if err := config.DB.Find(&admin).Error; err != nil {
		return nil, err
	}
	return admin, nil
}

func (r *AdminRepositoryStruct) FindByID(admin *models.Admin, id uint) error {
	return config.DB.First(admin, id).Error
}

func (r *AdminRepositoryStruct) Update(admin *models.Admin) error {
	if err := config.DB.Save(admin).Error; err != nil{
		return err
	}
	return nil
}

func (R *AdminRepositoryStruct) Delete(id uint) error {
	return config.DB.Delete(&models.Admin{}, id).Error
}