package db

import (
	"gorm.io/gorm"
)

// Owner is a struct to represent a user of the application
// Owner have a one car
// Owner have a name,last name email and relationship to car
type Owner struct {
	gorm.Model
	Name     string
	LastName string
	Email    string
	Car      Car `gorm:"foreignKey:ID"`
}
