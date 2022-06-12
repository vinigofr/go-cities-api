package service_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vinigofr/go-cities-api/internal/cities"
	"github.com/vinigofr/go-cities-api/internal/cities/mocks"
)

func TestGetAll(t *testing.T) {

    t.Run("Success to GetAll", func(t *testing.T) {
        mockedRepository := new(mocks.CityRepository)

        fakeCity := cities.City{
            Id:   1,
            Name: "Nova Ipixuna",
            UF:   "PA",
        }

        cityList := make([]cities.City, 0)
        cityList = append(cityList, fakeCity)
        mockedRepository.On("GetAll").Return(cityList, nil)

        service := cities.NewService(mockedRepository)

        result, err := service.GetAll()
        assert.NoError(t, err)

        assert.Equal(t, 1, result[0].Id)
        assert.Equal(t, "Nova Ipixuna", result[0].Name)
        assert.Equal(t, "PA", result[0].UF)
    })

    t.Run("Error case", func(t *testing.T) {
        expectedErr :=  errors.New("cannot get all cities")
        mockedRepository := new(mocks.CityRepository)

        mockedRepository.On("GetAll").
            Return(nil, expectedErr).
            Once()

        service := cities.NewService(mockedRepository)
        _, err := service.GetAll()
        assert.NotNil(t, err)

        mockedRepository.AssertExpectations(t)
    })
}
