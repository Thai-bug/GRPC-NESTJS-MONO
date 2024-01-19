package ResponseError

import (
	"encoding/json"
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

func ParseError(err error) []string {
	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		errorMessages := make([]string, len(validationErrs))
		for i, e := range validationErrs {
			println(e.Tag())
			switch e.Tag() {
			case "required_without":
				errorMessages[i] = fmt.Sprintf("The field %s is required if %s is not supplied", e.Field(), e.Param())
				break

			case "required":
				errorMessages[i] = fmt.Sprintf("The field %s is required", e.Field())
				break
			}
		}
		return errorMessages
	} else if marshallingErr, ok := err.(*json.UnmarshalTypeError); ok {
		return []string{fmt.Sprintf("The field %s must be a %s", marshallingErr.Field, marshallingErr.Type.String())}
	}
	return nil
}
