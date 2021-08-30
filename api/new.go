package api

import (
	db "cars/db"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string `json:"message"`
}

func NewCar(ctx *gin.Context) {

	car := db.Car{}
	ctx.ShouldBindJSON(&car)
	db.DB.Create(&car)
	ctx.JSON(200, car)
}

func NewOwner(ctx *gin.Context) {

	owner := db.Owner{}
	ctx.ShouldBindJSON(&owner)
	db.DB.Create(&owner)
	ctx.JSON(200, owner)
}
func NewBrand(ctx *gin.Context) {

	brand := db.Brand{}
	ctx.ShouldBindJSON(&brand)
	db.DB.Create(&brand)
	ctx.JSON(200, "Auto "+brand.Name+" creado")
}
