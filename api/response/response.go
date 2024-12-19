package resp

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

const (
	StatusOk    = "OK"
	StatusError = "Error"
)

var ErrorTokenExpired = errors.New("token expired")

func Error(msg string) Response {
	return Response{Status: StatusError, Error: msg}
}

func Ok() Response {
	return Response{Status: StatusOk}
}

func ValidateErrors(errs validator.ValidationErrors) Response {
	var errMassages []string

	for _, err := range errs {
		switch err.Tag() {
		case "required":
			errMassages = append(errMassages, fmt.Sprintf("%s is required", err.Field()))
		case "email":
			errMassages = append(errMassages, fmt.Sprintf("%s is not a valid email", err.Field()))
		}

	}
	return Error(strings.Join(errMassages, "; "))
}
