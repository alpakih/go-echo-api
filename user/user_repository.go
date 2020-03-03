package user

import (
	 "go-echo-api/entity"
)

type Repository interface {
	FindAll() ([]entity.User, error)
	FindById(id string) (entity.User, error)
	Save(dto Dto) (entity.User, error)
	Update(id string, dto Dto) (entity.User, error)
	Delete(id string) error
	Count() (uint, error)
}
