package mappers

import (
	"wells-go/application/dtos"
	"wells-go/domain/entities"
)

func ToRouteAccessResponseList(entities []entities.RouteAccessEntities) []dtos.RouteAccessResponse {
	var res []dtos.RouteAccessResponse
	for _, e := range entities {
		res = append(res, dtos.RouteAccessResponse{
			ID:             e.ID.String(),
			RoutePath:      e.RoutePath,
			HTTPMethod:     e.HTTPMethod,
			RoleName:       e.RoleName,
			PermissionName: e.PermissionName,
		})
	}
	return res
}
