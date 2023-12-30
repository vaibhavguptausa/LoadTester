package main

import "github.com/go-playground/validator/v10"

type CustomValidator struct {
	validator *validator.Validate
}

// Validate method for CustomValidator
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
