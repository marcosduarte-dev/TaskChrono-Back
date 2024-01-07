package database

import (
	"github.com/marcosduarte-dev/TaskChrono-Back/internal/entity"
	"gorm.io/gorm"
)

type Task struct {
	DB *gorm.DB
}

func NewTask(db *gorm.DB) *Task {
	return &Task{DB: db}
}

func (p *Task) Create(task *entity.Task) error {
	return p.DB.Create(task).Error
}

func (p *Task) FindMyTasks(page, limit int, sort string) ([]*entity.Task, error) {
	var task []*entity.Task
	var err error
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	if page != 0 && limit != 0 {
		err = p.DB.Limit(limit).Offset((page - 1) * limit).Order("project_id " + sort).Find(&task).Error
	} else {
		err = p.DB.Order("project_id " + sort).Preload("Project").Find(&task).Error
	}
	return task, err
}

func (p *Task) FindByProjectID(projectID string) ([]*entity.Task, error) {
	var task []*entity.Task

	err := p.DB.Preload("Project").Where("project_id = ?", projectID).Find(&task).Error

	return task, err
}

func (p *Task) FindByID(id string) (*entity.Task, error) {
	var task entity.Task
	err := p.DB.Preload("Project").First(&task, "id = ?", id).Error
	return &task, err
}

func (p *Task) Update(task *entity.Task) error {
	_, err := p.FindByID(task.ID.String())
	if err != nil {
		return err
	}
	return p.DB.Save(task).Error
}

func (p *Task) Delete(id string) error {
	task, err := p.FindByID(id)
	if err != nil {
		return err
	}
	return p.DB.Delete(task).Error
}