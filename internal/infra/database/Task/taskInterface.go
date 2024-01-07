package database

import "github.com/marcosduarte-dev/TaskChrono-Back/internal/entity"

type TaskInterface interface {
	Create(p *entity.Task) error
	FindMyTasks(page, limit int, sort string) ([]*entity.Task, error)
	FindByProjectID(rojectID string) ([]*entity.Task, error)
	FindByID(id string) (*entity.Task, error)
	Update(task *entity.Task) error
	Delete(id string) error
}