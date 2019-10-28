package model

import (
	"fmt"
	"net/http"
)

// Error contains error details for gRPC response and logging
type Error struct {
	Message    string `json:"message"`
	Status     uint32 `json:"-"`
	messageLog string
}

// NewErrorNil creates empty Error
func NewErrorNil() Error {
	return Error{}
}

// NewErrorPrecondition creates error with HTTP code 412
func NewErrorPrecondition(m string) Error {
	return newError(m, http.StatusPreconditionFailed)
}

// NewErrorServer creates error with HTTP code 500
func NewErrorServer(m string) Error {
	return newError(m, http.StatusInternalServerError)
}

// GetLogMessage returns message for log
func (e Error) GetLogMessage() string {
	if e.messageLog != "" {
		return e.messageLog
	}
	return e.Message
}

// IsError checks if Error contains error data
func (e Error) IsError() bool {
	return e.Status > 1
}

// WithError adds Golang error message to Error log message
func (e Error) WithError(err error) Error {
	e.messageLog = e.Message + ": " + err.Error()
	return e
}

// WithErrorMessage adds message to Error log message
func (e Error) WithErrorMessage(msg string) Error {
	e.messageLog = fmt.Sprintf("%s with message: %s", e.Message, msg)
	return e
}

// WithErrorObject copy log message from past Error
func (e Error) WithErrorObject(err Error) Error {
	e.messageLog = err.messageLog
	return e
}

// newError creates Error struct
func newError(m string, s uint32) Error {
	return Error{
		Message: m,
		Status:  s,
	}
}
