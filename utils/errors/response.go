package ResponseError

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func MapValidationErrors(err error) map[string]string {
	if errs, ok := err.(validator.ValidationErrors); ok {
		customErrors := make(map[string]string)

		for _, e := range errs {
			customErrors[e.Field()] = fmt.Sprintf("Invalid value for field %s", e.Field())
		}

		return customErrors
	}
	return nil
}
