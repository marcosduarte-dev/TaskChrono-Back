package database

import "github.com/marcosduarte-dev/TaskChrono-Back/internal/entity"

type ProjectInterface interface {
	Create(p *entity.Project) error
	FindMyProjects(page, limit int, sort string) ([]*entity.Project, error)
	FindByID(id string) (*entity.Project, error)
	Update(product *entity.Project) error
	Delete(id string) error
}