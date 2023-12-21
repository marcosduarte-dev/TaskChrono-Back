package dto

type ProjectInputDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
	UserID      string `json:"user_id"`
}