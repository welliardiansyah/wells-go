package mappers

import (
	"wells-go/application/dtos"
	"wells-go/domain/entities"
)

func ToPermissionResponse(permission entities.PermissionEntity) dtos.PermissionResponse {
	return dtos.PermissionResponse{
		ID:        permission.ID.String(),
		Name:      permission.Name,
		CanCreate: permission.CanCreate,
		CanRead:   permission.CanRead,
		CanUpdate: permission.CanUpdate,
		CanDelete: permission.CanDelete,
		CanExport: permission.CanExport,
		CanImport: permission.CanImport,
		CanView:   permission.CanView,
	}
}

func ToPermissionResponseList(permissions []entities.PermissionEntity) []dtos.PermissionResponse {
	responses := make([]dtos.PermissionResponse, 0, len(permissions))
	for _, p := range permissions {
		responses = append(responses, ToPermissionResponse(p))
	}
	return responses
}
