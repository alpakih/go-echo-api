package usecase

import (
	"github.com/jinzhu/gorm"
	"go-echo-api/auth"
	"go-echo-api/entity"
	"go-echo-api/utils"
)

type AuthService struct {
	*gorm.DB
}

func NewAuthService(db *gorm.DB) auth.Repository {
	return AuthService{db}
}

func (a AuthService) Login(email string) (entity.User, error) {
	var model entity.User
	err := a.DB.Find(&model, "email=?", email).Error
	return model, err
}

func (a AuthService) Register(dto auth.RegisterDto) (entity.User, error) {
	var model entity.User
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
