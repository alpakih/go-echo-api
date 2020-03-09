package auth

import "go-echo-api/models"

type Repository interface {
	Login(email string) (models.User, error)
	Register(dto RegisterDto) (models.User, error)
}
