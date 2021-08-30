package main

import (
	api "cars/api"
	"cars/db"

	"github.com/gin-gonic/gin"
)

func main() {

	db.Connect()
	server := gin.Default()
	server.GET("/cars", api.GetAllCars)
	server.GET("/carsbycolor", api.GetCarsByColor)
	server.POST("/cars", api.NewCar)
	server.POST("/owners", api.NewOwner)
	server.POST("/brands", api.NewBrand)
	server.Run()

}
