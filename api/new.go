package api

import (
	db "../db"
	"github.com/gin-gonic/gin"
)

func NewCar(ctx *gin.Context) {

	car := db.Car{}
	ctx.ShouldBindJSON(&car)
	db.DB.Create(&car)
	ctx.JSON(200, car)
}
