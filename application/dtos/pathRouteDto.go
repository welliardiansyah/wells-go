package dtos

import "time"

type PathRouteResponseDTO struct {
	ID        string    `json:"id"`         // snake_case
	Name      string    `json:"name"`       // snake_case
	Path      string    `json:"path"`       // snake_case
	Method    string    `json:"method"`     // snake_case
	CreatedAt time.Time `json:"created_at"` // snake_case
	UpdatedAt time.Time `json:"updated_at"` // snake_case
}

type PathRouteRequestDTO struct {
	Path   string `json:"path" binding:"required"`
	Method string `json:"method" binding:"required"`
}
