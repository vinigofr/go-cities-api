package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/vinigofr/go-cities-api/internal/cities"
)

type CityService struct {
	mock.Mock
}

func(p *CityService) Create() (cities.City, error) {
	args := p.Called()

	var city cities.City

	if rf, ok := args.Get(0).(func() cities.City); ok {
		city = rf()
	} else {
		if args.Get(0) != nil {
			city = args.Get(0).(cities.City)
		}
	}

	var err error

	if rf, ok := args.Get(1).(func() error); ok {
		err = rf()
	} else {
		if args.Get(1) != nil {
			err = args.Error(1)
		}
	}

	return city, err
}

func(p *CityService) GetAll() ([]cities.City, error) {
	args := p.Called()

	var allCities []cities.City

	if rf, ok := args.Get(0).(func() []cities.City); ok {
		allCities = rf()
	} else {
		if args.Get(0) != nil {
			allCities = args.Get(0).([]cities.City)
		}
	}

	var err error

	if rf, ok := args.Get(1).(func() error); ok {
		err = rf()
	} else {
		if args.Get(1) != nil {
			err = args.Error(1)
		}
	}

	return allCities, err
}