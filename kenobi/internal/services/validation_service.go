package services

import "github.com/go-playground/validator/v10"

var customValidator *validator.Validate

func init() {
	// TODO: get proper error message
	customValidator = validator.New()
}

// ValidateStruct validates a struct with tag style validation for that struct
func ValidateStruct(s any) error {
	return customValidator.Struct(s)
}

// ValidateVar validates a single variable using tag style validation
func ValidateVar(v any, tag string) error {
	return customValidator.Var(v, tag)
}
