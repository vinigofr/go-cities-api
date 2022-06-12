package cities

import (
	"errors"
)

type Service interface {
	Create(name, uf string) (City, error)
	GetAll() ([]City, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Create(name, uf string) (City, error) {
	city, err := s.repository.Create(name, uf)
	if err != nil {
		return City{}, errors.New("cannot create a city")
	}
	return city, nil
}

func (s *service) GetAll() ([]City, error) {
	cities, err := s.repository.GetAll()
	if err != nil {
		return nil,  errors.New("cannot get all cities")
	}

	return cities, nil
}