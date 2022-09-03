package apperrors

import (
	"encoding/json"
	"fmt"
)

var (
	ErrNotFound = NewAppError(nil, "not found", "", "US-000003")
)

type AppError struct {
	Err        error  `json:"-"`
	Message    string `json:"message,omitempty"`
	DevMessage string `json:"devMessage,omitempty"`
	Code       string `json:"code,omitempty"`
}

func NewAppError(err error, message, devMessage, code string) *AppError {
	if err == nil {
		err = fmt.Errorf(message)
	}
	return &AppError{
		Err:        err,
		Message:    message,
		DevMessage: devMessage,
		Code:       code,
	}
}

func (e AppError) Error() string {
	return e.Message
}

func (e AppError) Unwrap() error {
	return e.Err
}

func (e AppError) Marshal() []byte {
	marshal, err := json.Marshal(e)
	if err != nil {
		return nil
	}

	return marshal
}

func systemError(err error) *AppError {
	return NewAppError(err, err.Error(), "system error", "US-000000")
}
