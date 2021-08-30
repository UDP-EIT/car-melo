package db

import (
	"gorm.io/gorm"
)

type Brand struct {
	gorm.Model
	ID     int
	Name   string `json:"name"`
	Origin string `json:"origin"`
}
