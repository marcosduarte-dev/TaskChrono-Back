package database

import (
	"github.com/marcosduarte-dev/TaskChrono-Back/internal/entity"
	"gorm.io/gorm"
)

type Project struct {
	DB *gorm.DB
}

func NewProject(db *gorm.DB) *Project {
	return &Project{DB: db}
}

func (p *Project) Create(project *entity.Project) error {
	return p.DB.Create(project).Error
}

func (p *Project) FindMyProjects(page, limit int, sort string) ([]*entity.Project, error) {
	var project []*entity.Project
	var err error
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	if page != 0 && limit != 0 {
		err = p.DB.Limit(limit).Offset((page - 1) * limit).Order("user_id " + sort).Find(&project).Error
	} else {
		err = p.DB.Order("user_id " + sort).Find(&project).Error
	}
	return project, err
}

func (p *Project) FindByID(id string) (*entity.Project, error) {
	var project entity.Project
	err := p.DB.First(&project, "id = ?", id).Error
	return &project, err
}

func (p *Project) Update(project *entity.Project) error {
	_, err := p.FindByID(project.ID.String())
	if err != nil {
		return err
	}
	return p.DB.Save(project).Error
}

func (p *Project) Delete(id string) error {
	project, err := p.FindByID(id)
	if err != nil {
		return err
	}
	return p.DB.Delete(project).Error
}