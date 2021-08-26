package db

import (
	"gorm.io/gorm"
)

type Brand struct {
	gorm.Model
	Name   string
	Origin string
}
