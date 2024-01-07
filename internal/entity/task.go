package entity

import (
	"github.com/marcosduarte-dev/TaskChrono-Back/pkg/entity"
	"github.com/marcosduarte-dev/TaskChrono-Back/pkg/errors"
)

type Task struct {
	ID    			entity.ID 		`gorm:"primaryKey" json:"id"`
	Name  			string    		`json:"name"`
	Color 			string    		`json:"color"`
	Description string  			`json:"description"`	
	ProjectID   string        `json:"project_id"`
	Project     Project       `gorm:"foreignKey:ProjectID" json:"project"`
}

func NewTask(name string, color string, description string, projectID string) (*Task, error) {
	task := &Task{
		ID: entity.NewID(),
		Name: name,
		Color: color,
		Description: description,
		ProjectID: projectID,
	}
	err := task.ValidateTask()
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (p *Task) ValidateTask() error {
	if p.ID.String() == "" {
		return errors.ErrIDIsRequired
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil { 
		return errors.ErrInvalidID
	}
	if p.Name == "" {
		return errors.ErrNameIsRequired
	}
	if p.Color == "" {
		return errors.ErrColorIsRequired
	}
	if p.Description == "" {
		return errors.ErrDescriptionIsRequired
	}
	return nil
}