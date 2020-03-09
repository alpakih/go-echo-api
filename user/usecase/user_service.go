package usecase

import (
	"github.com/jinzhu/gorm"
	"go-echo-api/models"
	"go-echo-api/user"
	"go-echo-api/utils"
)

type UserService struct {
	*gorm.DB
}

func NewUserService(db *gorm.DB) user.Repository {
	return UserService{db}
}

func (u UserService) FindAll() ([]models.User, error) {
	var model []models.User
	err := u.DB.Find(&model).Error
	return model, err
}

func (u UserService) FindById(id string) (*models.User, error) {
	var model models.User
	model.ID = id
	err := u.DB.First(&model).Error
	if err != nil {
		return nil, err
	}
	return &model, err
}

func (u UserService) Save(dto user.Dto) (models.User, error) {
	var model models.User
	model.Name = dto.Name
	model.Email = dto.Email
	hashPassword, err := utils.HashPassword(dto.Password)
	if err != nil {
		return model, err
	}
	model.Password = hashPassword
	err = u.DB.Save(&model).Error
	return model, err
}

func (u UserService) Update(id string, updateDto user.Dto) (models.User, error) {
	var model models.User
	model.ID = id

	hashPassword, err := utils.HashPassword(updateDto.Password)
	if err != nil {
		return model, err
	}
	err = u.DB.Model(&model).UpdateColumns(models.User{
		Name:     updateDto.Name,
		Email:    updateDto.Email,
		Password: hashPassword,
	}).Error

	return model, err
}

func (u UserService) Delete(id string) (bool, error) {
	var model models.User
	model.ID = id
	isExisting, err := u.FindById(id)
	if isExisting == nil {
		return false, err
	}
	err = u.DB.Delete(&model).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
