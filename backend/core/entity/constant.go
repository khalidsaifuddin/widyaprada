package entity

import (
	"errors"
	"fmt"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrDuplicateKey   = errors.New("duplicate key violation")
	ErrInvalidData    = errors.New("invalid data")
	ErrDatabaseError  = errors.New("database error")
)

type RecordNotFoundError struct {
	Message string
	Err     error
}

func (e *RecordNotFoundError) Error() string {
	if e.Message != "" {
		return e.Message
	}
	return e.Err.Error()
}

func (e *RecordNotFoundError) Unwrap() error {
	return e.Err
}

func WrapRecordNotFound(message string) error {
	return &RecordNotFoundError{Message: message, Err: ErrRecordNotFound}
}

func WrapRecordNotFoundf(format string, args ...interface{}) error {
	return &RecordNotFoundError{Message: fmt.Sprintf(format, args...), Err: ErrRecordNotFound}
}

func IsRecordNotFound(err error) bool {
	var recordNotFoundErr *RecordNotFoundError
	if errors.As(err, &recordNotFoundErr) {
		return true
	}
	return errors.Is(err, ErrRecordNotFound)
}
