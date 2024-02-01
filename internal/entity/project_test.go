package entity

import (
	"testing"

	"github.com/marcosduarte-dev/TaskChrono-Back/pkg/errors"
	"github.com/stretchr/testify/assert"
)

const ProjectName = "Project 1"
const ProjectDescription = "Description"
const ProjectColor = "#eb4034"
const ProjectID = "1"

func TestNewProject(t *testing.T) {
	p, err := NewProject(ProjectName, ProjectColor, ProjectDescription, ProjectID)

	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.Equal(t, ProjectName, p.Name)
	assert.Equal(t, ProjectColor, p.Color)
	assert.Equal(t, ProjectDescription, p.Description)
}

func TestProjectWhenNameIsRequired(t *testing.T) {
	p, err := NewProject("", ProjectColor, ProjectDescription, ProjectID)

	assert.Nil(t, p)
	assert.NotNil(t, err)
	assert.Equal(t, errors.ErrNameIsRequired, err)
}

func TestProjectWhenColorIsRequired(t *testing.T) {
	p, err := NewProject(ProjectName, "", ProjectDescription, ProjectID)
	assert.Nil(t, p)
	assert.NotNil(t, err)
	assert.Equal(t, errors.ErrColorIsRequired, err)
}

func TestProjectWhenDescriptionIsRequired(t *testing.T) {
	p, err := NewProject(ProjectName, ProjectColor, "", ProjectID)
	assert.Nil(t, p)
	assert.NotNil(t, err)
	assert.Equal(t, errors.ErrDescriptionIsRequired, err)
}

func TestProjectValidate(t *testing.T) {
	p, err := NewProject(ProjectName, ProjectColor, ProjectDescription, ProjectID)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Nil(t, p.Validate())
}