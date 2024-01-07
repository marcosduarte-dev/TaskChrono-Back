package dto

type TaskInputDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
	ProjectID   string `json:"project_id"`
}