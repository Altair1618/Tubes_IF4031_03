package utils

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/valyala/fastjson"
)

type CustomValidator struct {
	Validator *validator.Validate `json:"validator"`
}

func (cv *CustomValidator) Validate(i interface{}) error {
	RegisterValidation(cv, "is_phone_number", PhoneNumberValidation)
	RegisterValidation(cv, "is_json", JsonValidation)
	RegisterValidation(cv, "is_date", DateValidation)
	RegisterValidation(cv, "is_division", DivisionValidation)
	RegisterValidation(cv, "is_category", CategoryValidation)
	RegisterValidation(cv, "is_sort", SortValidation)
	RegisterValidation(cv, "is_payment_status", PaymentStatusValidation)
	RegisterValidation(cv, "is_price", PriceValidation)
	RegisterValidation(cv, "is_seat_number", SeatNumberValidation)

	if err := cv.Validator.Struct(i); err != nil {
		return err
	}

	return nil
}

func RegisterValidation(cv *CustomValidator, validationName string, fn validator.Func) {
	err := cv.Validator.RegisterValidation(validationName, fn)

	if err != nil {
		fmt.Printf("Error registering %s validation", validationName)
	}
}

func GetValidationErrorMessages(err error) []FieldError {
	var errMessages []FieldError

	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				errMessages = append(errMessages, FieldError{Field: err.Field(), Message: fmt.Sprintf("%s is required", err.Field())})
			case "max":
				errMessages = append(errMessages, FieldError{Field: err.Field(), Message: fmt.Sprintf("%s maximum %s", err.Field(), err.Param())})
			case "min":
				errMessages = append(errMessages, FieldError{Field: err.Field(), Message: fmt.Sprintf("%s minimum %s", err.Field(), err.Param())})
			case "is_json":
				errMessages = append(errMessages, FieldError{Field: err.Field(), Message: fmt.Sprintf("%s not a valid json", err.Value())})
			case "is_date":
				errMessages = append(errMessages, FieldError{Field: err.Field(), Message: fmt.Sprintf("%s not a valid ISO 8601 date", err.Value())})
			case "is_sort":
				errMessages = append(errMessages, FieldError{Field: err.Field(), Message: fmt.Sprintf("%s not a valid sort", err.Value())})
			case "is_payment_status":
				errMessages = append(errMessages, FieldError{Field: err.Field(), Message: fmt.Sprintf("%s not a valid payment status", err.Value())})
			case "is_price":
				errMessages = append(errMessages, FieldError{Field: err.Field(), Message: "Price must be greater than or equal to 0"})
			case "is_seat_number":
				errMessages = append(errMessages, FieldError{Field: err.Field(), Message: fmt.Sprintf("%s not a valid seat number", err.Value())})
			}
		}
	}

	return errMessages
}

/* New validations */

func PhoneNumberValidation(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	if value == "" {
		return true
	}

	regex, _ := regexp.Compile(`^[0-9]{10,13}$`)

	return regex.MatchString(value)
}

func JsonValidation(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	if value == "" {
		return true
	}

	errParseJson := fastjson.Validate(value)

	return errParseJson == nil
}

func DateValidation(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	if value == "" {
		return true
	}

	regex, _ := regexp.Compile(`^(\d{4})-(\d{2})-(\d{2})T(\d{2}):(\d{2}):(\d{2}(?:\.\d*)?)((-(\d{2}):(\d{2})|Z)?)$`)

	return regex.MatchString(value)
}

func DivisionValidation(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	if value == "" {
		return true
	}

	return value == "ACADEMIC" || value == "COMPETITION"
}

func CategoryValidation(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	if value == "" {
		return true
	}

	return value == "PUBLIC" || value == "PRIVATE"
}

func SortValidation(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	if value == "" {
		return true
	}

	return value == "ASC" || value == "DESC"
}

func PaymentStatusValidation(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	if value == "" {
		return true
	}

	return value == "FAILED" || value == "SUCCESS"
}

func PriceValidation(fl validator.FieldLevel) bool {
	return fl.Field().Int() >= 0
}

func SeatNumberValidation(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	regex, _ := regexp.Compile(`^[A-Z]{2}[0-9]{3}$`)

	return regex.MatchString(value)
}
