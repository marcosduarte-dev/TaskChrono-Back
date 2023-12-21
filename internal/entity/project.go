package entity

import (
	"github.com/marcosduarte-dev/TaskChrono-Back/pkg/entity"
	"github.com/marcosduarte-dev/TaskChrono-Back/pkg/errors"
)

type Project struct {
	ID    			entity.ID 		`json:"id"`
	Name  			string    		`json:"name"`
	Color 			string    		`json:"color"`
	Description string  			`json:"description"`
	UserID      string 				`json:"user_id"`
}

func NewProject(name string, color string, description string, userID string) (*Project, error) {
	project := &Project{
		ID: entity.NewID(),
		Name: name,
		Color: color,
		Description: description,
		UserID: userID,
	}
	err := project.Validate()
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (p *Project) Validate() error {
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
	if p.UserID == "" {
		return errors.ErrUserIDIsRequired
	}
	return nil
}