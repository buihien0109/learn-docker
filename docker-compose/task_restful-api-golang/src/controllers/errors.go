package controllers

import (
	"errors"
	"github.com/sirupsen/logrus"
	"net/http"
)

var (
	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = errors.New("Internal Server Error ")
	// ErrNotFound will throw if the requested item is not exists
	ErrNotFound = errors.New("Your requested Item is not found ")
	// ErrConflict will throw if the current action already exists
	ErrConflict = errors.New("Your Item already exist ")
	// ErrBadParamInput will throw if the given request-body or params is not valid
	ErrBadParamInput = errors.New("Given Param is not valid ")
)

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case ErrInternalServerError:
		return http.StatusInternalServerError
	case ErrNotFound:
		return http.StatusNotFound
	case ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}