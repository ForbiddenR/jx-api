package rest

import "errors"

var (
	ErrBodyIsNil = errors.New("body is nil")
	// error return from services
	ErrServicesException = errors.New("services exception")
	// 404 not found error
	ErrNotFound = errors.New("not found")
)
