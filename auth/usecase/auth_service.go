package usecase

import (
	"github.com/jinzhu/gorm"
	"go-echo-api/auth"
	"go-echo-api/models"
	"go-echo-api/utils"
)

type AuthService struct {
	*gorm.DB
}

func NewAuthService(db *gorm.DB) auth.Repository {
	return AuthService{db}
}

func (a AuthService) Login(email string) (models.User, error) {
	var model models.User
	err := a.DB.Find(&model, "email=?", email).Error
	return model, err
}

func (a AuthService) Register(dto auth.RegisterDto) (models.User, error) {
	var model models.User
	model.Name = dto.Name
	model.Email = dto.Email
	hashPassword, err := utils.HashPassword(dto.Password)
	if err != nil {
		return model, err
	}
	model.Password = hashPassword
	err = a.DB.Save(&model).Error
	return model, err
}
