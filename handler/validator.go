package handler

import (
	validator "github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

// Validator ...
type Validator struct {
	validator *validator.Validate
}

// NewValidator ...
func NewValidator() echo.Validator {
	return &Validator{
		validator: validator.New(),
	}
}

// Validate validate request parameters
func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}
