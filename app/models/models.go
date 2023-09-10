package models

import "github.com/jinzhu/gorm"

type Person struct {
	gorm.Model
	Name	string	`json:"name" validate:"required"`
}