package mappers

import (
	"github.com/google/uuid"
	"wells-go/application/dtos"
	"wells-go/domain/entities"
)

func ToRouteAccessEntity(dto *dtos.RouteAccessRequestDTO) *entities.RouteAccessEntities {
	return &entities.RouteAccessEntities{
		ID:             uuid.New(),
		RoutePath:      dto.RoutePath,
		HTTPMethod:     dto.HTTPMethod,
		RoleName:       dto.RoleName,
		PermissionName: dto.PermissionName,
	}
}

func ToRouteAccessEntityWithID(id uuid.UUID, dto *dtos.RouteAccessRequestDTO) *entities.RouteAccessEntities {
	return &entities.RouteAccessEntities{
		ID:             id,
		RoutePath:      dto.RoutePath,
		HTTPMethod:     dto.HTTPMethod,
		RoleName:       dto.RoleName,
		PermissionName: dto.PermissionName,
	}
}
