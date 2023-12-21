package entity

import (
	"testing"

	"github.com/marcosduarte-dev/TaskChrono-Back/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestNewProject(t *testing.T) {
	p, err := NewProject("Project 1", "#eb4034", "Description", "1")

	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.Equal(t, "Project 1", p.Name)
	assert.Equal(t, "#eb4034", p.Color)
	assert.Equal(t, "Description", p.Description)
}

func TestProjectWhenNameIsRequired(t *testing.T) {
	p, err := NewProject("", "#eb4034", "Description", "1")

	assert.Nil(t, p)
	assert.NotNil(t, err)
	assert.Equal(t, errors.ErrNameIsRequired, err)
}

func TestProjectWhenColorIsRequired(t *testing.T) {
	p, err := NewProject("Project 1", "", "Description", "1")
	assert.Nil(t, p)
	assert.NotNil(t, err)
	assert.Equal(t, errors.ErrColorIsRequired, err)
}

func TestProjectWhenDescriptionIsRequired(t *testing.T) {
	p, err := NewProject("Project 1", "#eb4034", "", "1")
	assert.Nil(t, p)
	assert.NotNil(t, err)
	assert.Equal(t, errors.ErrDescriptionIsRequired, err)
}

func TestProjectValidate(t *testing.T) {
	p, err := NewProject("Project 1", "#eb4034", "Description", "1")
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Nil(t, p.Validate())
}