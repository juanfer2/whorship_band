package middlewares

import (
	"errors"

	"gorm.io/gorm"
)

type ErrorApi struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func NewErrorApi(message string, status int) *ErrorApi {
	return &ErrorApi{
		Message: message,
		Status:  status,
	}
}

type ErrorHandler struct {
	Error  ErrorApi
	Status int
}

func NewErrorHandler(err ErrorApi, status int) *ErrorHandler {
	return &ErrorHandler{
		Error:  err,
		Status: status,
	}
}

func ErrorHandlers(err error) *ErrorHandler {
	status := getStatus(err)
	errorApi := NewErrorApi(err.Error(), status)
	errorHandler := NewErrorHandler(*errorApi, status)

	return errorHandler
}

func getStatus(err error) int {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 404
	}

	if errors.Is(err, gorm.ErrInvalidField) || errors.Is(err, gorm.ErrUnsupportedRelation) ||
		errors.Is(err, gorm.ErrInvalidValue) {
		return 403
	}

	if errors.Is(err, gorm.ErrInvalidTransaction) {
		return 500
	}

	if errors.Is(err, gorm.ErrInvalidTransaction) {
		return 500
	}

	return 500
}
