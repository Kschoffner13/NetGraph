package types

import "net/http"

// ErrorResponse is the JSON shape returned to API clients on failure.
type ErrorResponse struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

// AppError is the internal error type handlers/services return. It carries
// the HTTP status it maps to, plus an optional wrapped error for logging.
type AppError struct {
	Status  int
	Code    string
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return e.Message + ": " + e.Err.Error()
	}
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

// Response converts the error to its JSON wire format.
func (e *AppError) Response() ErrorResponse {
	return ErrorResponse{Message: e.Message, Code: e.Code}
}

func NewBadRequest(message string) *AppError {
	return &AppError{Status: http.StatusBadRequest, Code: "bad_request", Message: message}
}

func NewUnauthorized(message string) *AppError {
	return &AppError{Status: http.StatusUnauthorized, Code: "unauthorized", Message: message}
}

func NewForbidden(message string) *AppError {
	return &AppError{Status: http.StatusForbidden, Code: "forbidden", Message: message}
}

func NewNotFound(message string) *AppError {
	return &AppError{Status: http.StatusNotFound, Code: "not_found", Message: message}
}

func NewConflict(message string) *AppError {
	return &AppError{Status: http.StatusConflict, Code: "conflict", Message: message}
}

// NewInternal wraps an unexpected error without leaking its details to the client.
func NewInternal(err error) *AppError {
	return &AppError{Status: http.StatusInternalServerError, Code: "internal_error", Message: "something went wrong", Err: err}
}

// NewUpstreamError wraps a failure from an external dependency (e.g. the ISED API).
func NewUpstreamError(err error) *AppError {
	return &AppError{Status: http.StatusBadGateway, Code: "upstream_error", Message: "upstream service failed", Err: err}
}
