package user

import "go-echo-api/entity"

type Mapper struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUserMapper() *Mapper {
	return &Mapper{}
}
func (m *Mapper) Map(model entity.User) *Mapper {
	m.ID = model.ID
	m.Name = model.Name
	m.Email = model.Email
	return m
}

func (m *Mapper) MapList(model []entity.User) interface{} {
	var length = len(model)
	serialized := make([]Mapper, length, length)

	for k, v := range model {
		serialized[k] = Mapper{
			ID:    v.ID,
			Name:  v.Name,
			Email: v.Email,
		}
	}
	return serialized
}
