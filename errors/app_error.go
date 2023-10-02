package errors

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Message     string `json:"message"`
	Code        string `json:"code"`
	Description string `json:"description"`
	HTTPStatus  int
}

func (e *AppError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// NotFoundError represents an error when a resource is not found.
type NotFoundError struct {
	*AppError
}

func NewNotFoundError(message, description string, code string) *NotFoundError {
	return &NotFoundError{
		AppError: &AppError{
			Message:     message,
			Code:        code,
			Description: description,
			HTTPStatus:  http.StatusNotFound,
		},
	}
}

type UnauthorizedError struct {
	*AppError
}

func NewUnauthorizedError(message, description string, code string) *UnauthorizedError {
	return &UnauthorizedError{
		AppError: &AppError{
			Message:     message,
			Code:        code,
			Description: description,
			HTTPStatus:  http.StatusUnauthorized,
		},
	}
}

type BadRequestError struct {
	*AppError
}

func NewBadRequestError(message, description string, code string) *BadRequestError {
	return &BadRequestError{
		AppError: &AppError{
			Message:     message,
			Code:        code,
			Description: description,
			HTTPStatus:  http.StatusBadRequest,
		},
	}
}
