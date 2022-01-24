package utils

import (
	"fmt"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

type ReturnErrors struct {
	Errors []ReturnError `json:"message"`
}

type ReturnError struct {
	Error string `json:"error"`
}

type AnyValidator struct {
	Validator *validator.Validate
}

func (av *AnyValidator) Validate(i interface{}) error {
	if err := av.Validator.Struct(i); err != nil {
		return err
	}
	return nil
}

func TranslateError(err error, validate *validator.Validate) (errs []error) {
	if err == nil {
		return nil
	}
	uni := ut.New(en.New(), en.New())
	trans, _ := uni.GetTranslator("en")
	_ = enTranslations.RegisterDefaultTranslations(validate, trans)
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(trans))
		errs = append(errs, translatedErr)
	}
	return errs
}

func ReturnErros(err error, validate AnyValidator) ReturnErrors {
	errors := TranslateError(err, validate.Validator)
	returnErrors := ReturnErrors{}
	for _, e := range errors {
		returnErrors.Errors = append(returnErrors.Errors, ReturnError{Error: e.Error()})
	}

	return returnErrors

}
