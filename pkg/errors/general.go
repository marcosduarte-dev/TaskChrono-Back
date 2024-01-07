package errors

import "errors"

var (
	ErrNameIsRequired        = errors.New("name is required")
	ErrColorIsRequired       = errors.New("color is required")
	ErrDescriptionIsRequired = errors.New("description is required")
	ErrUserIDIsRequired      = errors.New("UserID is required")
	ErrJSON                  = errors.New("JSON has problems")
)