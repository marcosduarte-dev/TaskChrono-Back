package entity

import (
	"time"

	"github.com/marcosduarte-dev/TaskChrono-Back/pkg/entity"
	"github.com/marcosduarte-dev/TaskChrono-Back/pkg/errors"
)

type Timer struct {
	ID    			   entity.ID 		 `gorm:"primaryKey" json:"id"`
	StartTime      time.Time     `json:"start_time"`
  EndTime        time.Time     `json:"end_time"`
  TotalDuration  int           `json:"total_duration"`
  RecordType     string        `json:"record_type"`
	TaskID         string        `json:"task_id"`
	Task           Task          `gorm:"foreignKey:TaskID" json:"task"`
}

func NewTimer(startTime time.Time, endTime time.Time, TotalDuration int, RecordType string, taskID string) (*Timer, error) {
	timer := &Timer{
		ID: entity.NewID(),
		StartTime: startTime,
		EndTime: endTime,
		TotalDuration: TotalDuration,
		RecordType: RecordType,
		TaskID: taskID,
	}
	err := timer.ValidateTimer()
	if err != nil {
		return nil, err
	}
	return timer, nil
}

func (t *Timer) ValidateTimer() error {
	if t.ID.String() == "" {
		return errors.ErrIDIsRequired
	}
	
	_, err := entity.ParseID(t.ID.String())
	if err != nil {
		return errors.ErrInvalidID
	}

	switch t.RecordType {
	case "start":
		if t.TotalDuration > 0 {
			return errors.ErrTotalDurationStart
		}
		if t.StartTime.IsZero() {
			return errors.ErrStartTimeIsRequired
		}
		if !t.EndTime.IsZero() {
			return errors.ErrEndTimeMustBeNull
		}
	case "stop":
		if t.TotalDuration <= 0 {
			return errors.ErrTotalDuration
		}
		if !t.StartTime.IsZero() {
			return errors.ErrStartTimeMustBeNull
		}
		if t.EndTime.IsZero() {
			return errors.ErrEndTimeIsRequired
		}
	default:
		return errors.ErrRecordTypeValidValues
	}

	if t.RecordType == "" {
		return errors.ErrRecordTypeIsRequired
	}

	return nil
}
