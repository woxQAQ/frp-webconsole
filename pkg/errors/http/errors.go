package http

import (
	"fmt"
	"net/http"
)

type ErrorHandler interface {
	JSON(code int, obj any)
}

type HttpError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Error(handler ErrorHandler, code int, message string) {
	handler.JSON(code, HttpError{
		Code:    code,
		Message: message,
	})
}

func (e *HttpError) Error() string {
	return fmt.Sprintf("HTTP %d: %s", e.Code, e.Message)
}

func NewBadGatewayError(handler ErrorHandler, message string) {
	handler.JSON(http.StatusBadGateway, HttpError{
		Code:    http.StatusBadGateway,
		Message: message,
	})
}

func NewBadRequestError(handler ErrorHandler, message string) {
	handler.JSON(http.StatusBadRequest, HttpError{
		Code:    http.StatusBadRequest,
		Message: message,
	})
}

func NewConflictError(handler ErrorHandler, message string) {
	handler.JSON(http.StatusConflict, HttpError{
		Code:    http.StatusConflict,
		Message: message,
	})
}

func NewForbiddenError(handler ErrorHandler, message string) {
	handler.JSON(http.StatusForbidden, HttpError{
		Code:    http.StatusForbidden,
		Message: message,
	})
}

func NewInternalServerError(handler ErrorHandler, message string) {
	handler.JSON(http.StatusInternalServerError, HttpError{
		Code:    http.StatusInternalServerError,
		Message: message,
	})
}

func NewNotFoundError(handler ErrorHandler, message string) {
	handler.JSON(http.StatusNotFound, HttpError{
		Code:    http.StatusNotFound,
		Message: message,
	})
}

func NewUnauthorizedError(handler ErrorHandler, message string) {
	handler.JSON(http.StatusUnauthorized, HttpError{
		Code:    http.StatusUnauthorized,
		Message: message,
	})
}

func NewUnprocessableEntityError(handler ErrorHandler, message string) {
	handler.JSON(http.StatusUnprocessableEntity, HttpError{
		Code:    http.StatusUnprocessableEntity,
		Message: message,
	})
}

func NewTooManyRequestsError(handler ErrorHandler, message string) {
	handler.JSON(http.StatusTooManyRequests, HttpError{
		Code:    http.StatusTooManyRequests,
		Message: message,
	})
}
