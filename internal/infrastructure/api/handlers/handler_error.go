package handlers

import (
	"errors"
	internalerrors "neoway_test/internal/internal-errors"
	"net/http"

	"github.com/go-chi/render"
	"gorm.io/gorm"
)

type EndpointFunc func(w http.ResponseWriter, r *http.Request) (interface{}, int, error)

type MultiError struct {
	Errors []error
}

func (e *MultiError) Error() string {
	var errMsg string
	for _, err := range e.Errors {
		errMsg += err.Error() + "; "
	}
	return errMsg
}

func HandlerError(endpointFunc EndpointFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		obj, status, err := endpointFunc(w, r)

		if err != nil && obj == nil {
			handleError(w, r, err)
			return
		}

		if err != nil && obj != nil {
			handleErrorAndSuccess(w, r, err, obj)
			return
		}

		render.Status(r, status)

		if err == nil && obj != nil {
			render.JSON(w, r, obj)
		}
	})
}

func handleError(w http.ResponseWriter, r *http.Request, err error) {
	var status int

	switch {
	case errors.Is(err, internalerrors.ErrInternal):
		status = 500
	case errors.Is(err, gorm.ErrRecordNotFound):
		status = 404
	default:
		status = 400
	}

	render.Status(r, status)

	// Use a type assertion to check if it's a MultiError
	if multiErr, ok := err.(*MultiError); ok {
		render.JSON(w, r, map[string]string{"error": multiErr.Error()})
	} else {
		render.JSON(w, r, map[string]string{"error": err.Error()})
	}
}

func handleErrorAndSuccess(w http.ResponseWriter, r *http.Request, err error, obj interface{}) {
	var status int

	switch {
	case errors.Is(err, internalerrors.ErrInternal):
		status = 500
	case errors.Is(err, gorm.ErrRecordNotFound):
		status = 404
	default:
		status = 201
	}

	render.Status(r, status)

	// Use a type assertion to check if it's a MultiError
	if multiErr, ok := err.(*MultiError); ok {
		render.JSON(w, r, map[string]interface{}{"res": obj, "alertMsg": multiErr.Error()})
	} else {
		render.JSON(w, r, map[string]interface{}{"res": obj, "alertMsg": err.Error()})
	}
}
