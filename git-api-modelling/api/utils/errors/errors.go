package errors

import (
	"encoding/json"
	"errors"
	"net/http"
)

// ApiError interface definition
type ApiError interface {
	Status() int
	Message() string
	Error() string
}

type apiError struct {
	AStatus  int    `json:"status"`
	AMessage string `json:"message"`
	AnError  string `json:"error,omitempty"`
}

func (e *apiError) Status() int {
	return e.AStatus
}

func (e *apiError) Message() string {
	return e.AMessage
}

func (e *apiError) Error() string {
	return e.AnError
}

// ApiError ...
func NewApiError(statusCode int, message string) ApiError {
	return &apiError{AStatus: statusCode, AMessage: message}
}

// NewApiErrFromBytes ...
func NewApiErrFromBytes(body []byte) (ApiError, error) {
	var result apiError
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.New("invalid json body")
	}
	return &result, nil
}

// NewInternalServerError ...
func NewInternalServerError(message string) ApiError {
	return &apiError{
		AStatus:  http.StatusInternalServerError,
		AMessage: message,
	}
}

// NewNotFoundError ...
func NewNotFoundError(message string) ApiError {
	return &apiError{
		AStatus:  http.StatusNotFound,
		AMessage: message,
	}
}

// NewBadRequestError ...
func NewBadRequestError(message string) ApiError {
	return &apiError{
		AStatus:  http.StatusBadRequest,
		AMessage: message,
	}
}
