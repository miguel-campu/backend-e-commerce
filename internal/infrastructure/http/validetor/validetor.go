package validator

import (
	v "github.com/go-playground/validator/v10" // alias 'v'
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *v.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func NewValidator() echo.Validator {
	return &CustomValidator{validator: v.New()}
}
