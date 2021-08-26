package api

import (
	db "../db"
	"github.com/gin-gonic/gin"
)

func GetAllCars(ctx *gin.Context) {

	cars := []db.Car{}

	db.DB.Find(&cars)
	ctx.JSON(200, cars)
}

func GetCarsByColor(ctx *gin.Context) {

	color := ctx.Request.URL.Query().Get("color")
	car := db.Car{}
	db.DB.Where("color = ?", color).First(&car)
	ctx.JSON(200, car)
}
