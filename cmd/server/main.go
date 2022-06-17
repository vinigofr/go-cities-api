package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinigofr/go-cities-api/cmd/server/controllers"
	"github.com/vinigofr/go-cities-api/internal/cities"
)



func main() {
	server := gin.Default()

	citiesRepo := cities.NewRepository()
	citiesService := cities.NewService(citiesRepo)
	citiesController := controllers.NewCity(citiesService)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "pong")
	})

	server.GET("/city", citiesController.GetAll())
	server.POST("/city", citiesController.Create())
	
	server.Run(":4000")


}
