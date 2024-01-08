package dto

import "time"

type TimerInputDTO struct {
	StartTime     time.Time `json:"start_time"`
	EndTime       time.Time `json:"end_time"`
	TotalDuration int       `json:"total_duration"`
	RecordType    string    `json:"record_type"`
	TaskID        string    `json:"task_id"`
}