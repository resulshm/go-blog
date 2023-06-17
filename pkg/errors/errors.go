package errors

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/resulshm/go-blog/internal/providers/validation"
)

var errorList = make(map[string]string)

func Get() map[string]string {
	return errorList
}

func Init() {
	errorList = map[string]string{}
}

func SetFromErrors(err error) {
	var validationErrors validator.ValidationErrors

	if errors.As(err, &validationErrors) {
		for _, fieldError := range validationErrors {
			Add(fieldError.Field(), GetErrorMsg(fieldError.Tag()))
		}
	}
}

func Add(key, value string) {
	errorList[strings.ToLower(key)] = value
}

func GetErrorMsg(tag string) string {
	return validation.ErrorMessage()[tag]
}
