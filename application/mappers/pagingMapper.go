package mappers

import "wells-go/application/dtos"

func ToPagingResponseFlat[T any](items []T, page, limit int, total int64) dtos.PagingResponseFlat[T] {
	totalPages := int64(0)
	if limit > 0 {
		totalPages = (total + int64(limit) - 1) / int64(limit)
	}
	return dtos.PagingResponseFlat[T]{
		Data:       items,
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: totalPages,
	}
}
