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

func (p *Timer) ValidateTimer() error {
	print(p.StartTime.String())
	if p.ID.String() == "" {
		return errors.ErrIDIsRequired
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil { 
		return errors.ErrInvalidID
	}
	if p.RecordType == "start" {
		if p.TotalDuration > 0 {
			return errors.ErrTotalDurationStart
		}
		if p.StartTime.IsZero() {
			return errors.ErrStartTimeIsRequired
		}
		if !p.EndTime.IsZero() {
			return errors.ErrEndTimeMustBeNull
		}
	} else {
		if p.TotalDuration <= 0 {
			return errors.ErrTotalDuration
		}
		if !p.StartTime.IsZero() {
			return errors.ErrStartTimeMustBeNull
		}
		if p.EndTime.IsZero() {
			return errors.ErrEndTimeIsRequired
		}
	}
	if p.RecordType == "" {
		return errors.ErrRecordTypeIsRequired
	}
	if p.RecordType != "start" && p.RecordType != "stop" {
		return errors.ErrRecordTypeValidValues
	}
	return nil
}