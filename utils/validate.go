package utils

import (
	"simple-crud-app/app/models"

	"github.com/go-playground/validator/v10"
)


func InputValidator(p models.Person) error {
	validate := validator.New()

	err := validate.Struct(p)
	if err != nil {
		return err
	}
	return nil
}