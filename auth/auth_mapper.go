package auth

import "go-echo-api/models"

type Mapper struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewAuthMapper() *Mapper {
	return &Mapper{}
}
func (m *Mapper) Map(model models.User) *Mapper {
	m.ID = model.ID
	m.Name = model.Name
	m.Email = model.Email
	return m
}

