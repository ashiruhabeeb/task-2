package utils

import (
	"simple-crud-app/api/models"

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

func InputPayloadValidator(p models.Payload) error {
	validate := validator.New()

	err := validate.Struct(p)
	if err != nil {
		return err
	}
	return nil
}