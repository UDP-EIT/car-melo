package db

import (
	"gorm.io/gorm"
)

type Car struct {
	gorm.Model
	Color  string `json:"color"`
	Modelo string `json:"model"`
	Year   int    `json:"year"`
	Price  int    `json:"price"`
	Engine string `json:"engine"`
	Brand  Brand  `gorm:"foreignKey:ID"`
}
