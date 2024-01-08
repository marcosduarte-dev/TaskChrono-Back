package errors

import "errors"

var (
	ErrTimerNotFound       				 = errors.New("timer not found")
	ErrTimerUpdate         				 = errors.New("error on timer update")
	ErrTimerDelete         				 = errors.New("error on task delete")
	ErrRecordTypeIsRequired        = errors.New("record type is required")
	ErrRecordTypeValidValues       = errors.New("record type must be 'start' or 'stop'")
	ErrStartTimeIsRequired         = errors.New("start time is required")
	ErrStartTimeMustBeNull         = errors.New("start time must be null")
	ErrEndTimeIsRequired           = errors.New("end time is required")
	ErrEndTimeMustBeNull           = errors.New("end time must be null")
	ErrTotalDuration               = errors.New("total duration must be greater than 0")
	ErrTotalDurationStart          = errors.New("total duration on start must be 0")
)
