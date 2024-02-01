package database

import (
	"time"

	"github.com/marcosduarte-dev/TaskChrono-Back/internal/entity"
	"gorm.io/gorm"
)

type Timer struct {
	DB *gorm.DB
}

const PreLoadTaskProject = "Task.Project"

func NewTimer(db *gorm.DB) *Timer {
	return &Timer{DB: db}
}

func (p *Timer) Create(timer *entity.Timer) error {
	return p.DB.Create(timer).Error
}

func (p *Timer) FindMyTimers(page, limit int, sort string) ([]*entity.Timer, error) {
	var timer []*entity.Timer
	var err error
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	if page != 0 && limit != 0 {
		err = p.DB.Limit(limit).Offset((page - 1) * limit).Order("task_id " + sort).Find(&timer).Error
	} else {
		err = p.DB.Order("task_id " + sort).Preload("Task").Preload(PreLoadTaskProject).Find(&timer).Error
	}
	return timer, err
}

func (p *Timer) FindByTaskID(taskID string) ([]*entity.Timer, error) {
	var timer []*entity.Timer

	err := p.DB.Preload("Task").Preload(PreLoadTaskProject).Where("task_id = ?", taskID).Find(&timer).Error

	return timer, err
}

func (p *Timer) FindByID(id string) (*entity.Timer, error) {
	var timer entity.Timer
	err := p.DB.Preload("Task").First(&timer, "id = ?", id).Error
	return &timer, err
}

func (p *Timer) FindByDate(date time.Time) ([]*entity.Timer, error) {
	var timer []*entity.Timer

	err := p.DB.Where("DATE(start_time) = ? OR DATE(end_time) = ?", date.Format("2006-01-02"), date.Format("2006-01-02")).Preload("Task").Preload(PreLoadTaskProject).Find(&timer).Error

	return timer, err
}

func (p *Timer) Update(timer *entity.Timer) error {
	_, err := p.FindByID(timer.ID.String())
	if err != nil {
		return err
	}
	return p.DB.Save(timer).Error
}

func (p *Timer) Delete(id string) error {
	timer, err := p.FindByID(id)
	if err != nil {
		return err
	}
	return p.DB.Delete(timer).Error
}