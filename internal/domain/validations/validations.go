package validations

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(obj interface{}) []error {

	errs := make([]error, 0)

	validate := validator.New()

	err := validate.Struct(obj)

	validationErros := err.(validator.ValidationErrors)

	for _, v := range validationErros {

		switch v.Tag() {
		case "required":
			errs = append(errs, errors.New(v.Field()+"is required"))
		case "max":
			errs = append(errs, errors.New(v.Field()+" exceeded max lengh"+v.Param()))
		case "min":
			errs = append(errs, errors.New(v.Field()+" mininum lengh is required"+v.Param()))
		case "email":
			errs = append(errs, errors.New(v.Field()+" is invalid"+v.Param()))
		}
	}

	if errs != nil {
		return errs
	}

	return nil
}
