package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/vinigofr/go-cities-api/internal/cities"
)

type CityRepository struct {
	mock.Mock
}

func(c *CityRepository) GetAll() ([]cities.City, error) {
	// To success
	args := c.Called()

	var city []cities.City

	if rf, ok := args.Get(0).(func() []cities.City); ok {
		city = rf()
	} else {
		if args.Get(0) != nil {
			city = args.Get(0).([]cities.City)
		}
	}

	// To error
	var err error

	if rf, ok := args.Get(1).(func() error); ok {
		err = rf()
	} else {
		err = args.Error(1)
	}

	return city, err
}

func(c *CityRepository) Create(name, uf string) (cities.City, error) {
	args := c.Called()

	var city cities.City

	if rf, ok := args.Get(0).(func(name, uf string) cities.City); ok {
		city = rf(name, uf)
	} else {
		if args.Get(0) != nil {
			city = args.Get(0).(cities.City)
		}
	}

	var err error

	if rf, ok := args.Get(1).(func() error); ok {
		err = rf()
	} else {
		err = args.Error(1)
	}

	return city, err
}