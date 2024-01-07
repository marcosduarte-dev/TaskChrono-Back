package errors

import "errors"

var (
	ErrProjectNotFound       = errors.New("project not found")
	ErrProjectUpdate         = errors.New("error on project update")
	ErrProjectDelete         = errors.New("error on project delete")
)
