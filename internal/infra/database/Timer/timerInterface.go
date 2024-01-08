package database

import "github.com/marcosduarte-dev/TaskChrono-Back/internal/entity"

type TimerInterface interface {
	Create(p *entity.Timer) error
	FindMyTimers(page, limit int, sort string) ([]*entity.Timer, error)
	FindByTaskID(taskID string) ([]*entity.Timer, error)
	FindByID(id string) (*entity.Timer, error)
	Update(timer *entity.Timer) error
	Delete(id string) error
}