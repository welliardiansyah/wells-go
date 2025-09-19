package dtos

type PagingResponseFlat[T any] struct {
	Data       []T   `json:"data"` // data array
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int64 `json:"totalPages"`
}
