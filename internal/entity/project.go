package entity

import (
	"errors"

	"github.com/marcosduarte-dev/TaskChrono-Back/pkg/entity"
)

var (
	ErrIDIsRequired = errors.New("id is required")
	ErrInvalidID = errors.New("invalid id")
	ErrNameIsRequired = errors.New("name is required")
	ErrColorIsRequired = errors.New("color is required")
	ErrDescriptionIsRequired = errors.New("color is required")
)

type Project struct {
	ID    			entity.ID 		`json:"id"`
	Name  			string    		`json:"name"`
	Color 			string    		`json:"color"`
	Description string  			`json:"description"`
}

func NewProject(name string, color string, description string) (*Project, error) {
	project := &Project{
		ID: entity.NewID(),
		Name: name,
		Color: color,
		Description: description,
	}
	err := project.Validate()
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (p *Project) Validate() error {
	if p.ID.String() == "" {
		return ErrIDIsRequired
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil { 
		return ErrInvalidID
	}
	if p.Name == "" {
		return ErrNameIsRequired
	}
	if p.Color == "" {
		return ErrColorIsRequired
	}
	if p.Description == "" {
		return ErrDescriptionIsRequired
	}
	return nil
}