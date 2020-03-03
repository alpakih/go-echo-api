package auth

import "go-echo-api/entity"

type Repository interface {
	Login(email string) (entity.User, error)
	Register(dto RegisterDto) (entity.User, error)
}
