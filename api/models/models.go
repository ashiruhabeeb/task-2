package models

import "github.com/google/uuid"

type Person struct {
	ID			uuid.UUID	`gorm:"type:uuid;primary_key" json:"id"`
	Name		string		`json:"name" gorm:"unique" validate:"required,min=2"`
}

type Payload struct {
	Name		string		`json:"name" gorm:"unique" validate:"required,min=2"`
}