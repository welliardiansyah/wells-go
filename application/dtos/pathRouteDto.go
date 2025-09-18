package dtos

import "time"

type PathRouteResponseDTO struct {
	ID        string    `json:"id"`
	Path      string    `json:"path"`
	Method    string    `json:"method"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PathRouteRequestDTO struct {
	Path   string `json:"path" binding:"required"`
	Method string `json:"method" binding:"required"`
}
