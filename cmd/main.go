package main

import (
	"../db"
	"github.com/gin-gonic/gin"

	api "../api"
)

func main() {

	db.Connect()
	server := gin.Default()
	server.GET("/cars", api.GetAllCars)
	server.GET("/carsbycolor", api.GetCarsByColor)
	server.POST("/cars", api.NewCar)
	server.Run()

}
