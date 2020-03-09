package user

import (
	"go-echo-api/models"
)

type Repository interface {
	FindAll() ([]models.User, error)
	FindById(id string) (*models.User, error)
	Save(dto Dto) (models.User, error)
	Update(id string, dto Dto) (models.User, error)
	Delete(id string) (bool, error)
}
