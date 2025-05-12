package helper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type ValidationError map[string]string

func (v ValidationError) Error() string {
	return "validation error"
}

func ValidateStruct(s interface{}) error {
	if err := validate.Struct(s); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			errorsMap := make(map[string]string)
			for _, e := range ve {
				errorsMap[e.Field()] = fmt.Sprintf("Error on field %s, condition must be %s", e.Field(), e.Tag())
			}
			return ValidationError(errorsMap)
		}
		return err
	}
	return nil
}