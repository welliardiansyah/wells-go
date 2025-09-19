package mappers

import (
	"wells-go/application/dtos"
	"wells-go/domain/entities"
)

func ToPathRouteResponse(entity *entities.PathRouteEntities) *dtos.PathRouteResponseDTO {
	if entity == nil {
		return nil
	}
	return &dtos.PathRouteResponseDTO{
		ID:        entity.ID.String(),
		Path:      entity.Path,
		Name:      entity.Name,
		Method:    entity.Method,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}

func ToPathRouteResponseList(entities []*entities.PathRouteEntities) []*dtos.PathRouteResponseDTO {
	res := make([]*dtos.PathRouteResponseDTO, 0, len(entities))
	for _, e := range entities {
		res = append(res, ToPathRouteResponse(e))
	}
	return res
}
