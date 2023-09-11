package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Person struct {
	ID			int			`gorm:"primaryKey,autoIncrement" json:"id"`
	Name		string		`json:"name" gorm:"unique" validate:"required,min=2"`
	CreatedAt	time.Time	`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt	time.Time	`gorm:"defualt:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt	gorm.DeletedAt	`gorm:"index"`
}

func (p *Person) Validate() error {
	validate := validator.New()
	
	if err := validate.Struct(&p); err != nil {
		return err
	}
	return nil
}