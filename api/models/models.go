package models

type Person struct {
	ID		string	`gorm:"primary_key" json:"id"`
	Name	string	`json:"name" gorm:"unique" validate:"required,min=2"`
}

type Payload struct {
	Name		string		`json:"name" gorm:"unique" validate:"required,min=2"`
}