package repository

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vinigofr/go-cities-api/internal/cities"
)

func TestRepositoryGetAll(t *testing.T) {
	t.Run("should return a valid city list", func(t *testing.T) {
		fakeCities := []cities.City{
			{Id: 1, Name:  "Presidente Dutra", UF: "MA"},
			{Id: 2, Name:  "Tuntum", UF: "MA"},
			{Id: 3, Name:  "Dom Pedro", UF: "MA"},
			{Id: 4, Name:  "Barra do Corda", UF: "MA"},
		}

		cities.Database = append(cities.Database, fakeCities...)

		repository := cities.NewRepository()

		result, _ := repository.GetAll()

		byteResponse, _ := json.Marshal(fakeCities)
		byteExpected, _ := json.Marshal(result)

		assert.Equal(t, string(byteExpected), string(byteResponse))
		cities.Database = nil
	})
}

func TestRepositoryCreate(t *testing.T) {
	t.Run("should return a created data when success", func(t *testing.T) {
		fakeCities := []cities.City{
			{Id: 1, Name:  "Presidente Dutra", UF: "MA"},
			{Id: 2, Name:  "Tuntum", UF: "MA"},
			{Id: 3, Name:  "Dom Pedro", UF: "MA"},
			{Id: 4, Name:  "Barra do Corda", UF: "MA"},
		}
		
		repository := cities.NewRepository()
		for _, city := range fakeCities {
			result, _ := repository.Create(city.Name, city.UF)
			bData, _ := json.Marshal(result)
			bFakeCit, _ := json.Marshal(city)

			assert.Equal(t, string(bFakeCit), string(bData))
		}
	})
}

