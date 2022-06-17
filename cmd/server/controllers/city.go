package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinigofr/go-cities-api/internal/cities"
)

type CityController struct {
	cities.Service
}

type requestCity struct {
	Name string `json:"name" binding:"required"`
	UF string `json:"uf" binding:"required"`
}

func NewCity(s cities.Service) *CityController {
	return &CityController{
		Service: s,
	}
}

func (s *CityController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var requestData requestCity

		if err := ctx.ShouldBindJSON(&requestData); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "invalid body request",
			})
			return
		}

		city, err := s.Service.Create(requestData.Name, requestData.UF)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusCreated, city)
	}
}
func (s *CityController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cities, err := s.Service.GetAll()

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, cities)
	}
}