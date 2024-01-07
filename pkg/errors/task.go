package errors

import "errors"

var (
	ErrTaskNotFound       = errors.New("task not found")
	ErrTaskUpdate         = errors.New("error on task update")
	ErrTaskDelete         = errors.New("error on task delete")
)
