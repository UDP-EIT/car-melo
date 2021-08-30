package db

import (
	"gorm.io/gorm"
)

// Owner is a struct to represent a user of the application
// Owner have a one car
// Owner have a name,last name email and relationship to car
type Owner struct {
	gorm.Model
	Id       int
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	CarId    int
	Car      Car
}
