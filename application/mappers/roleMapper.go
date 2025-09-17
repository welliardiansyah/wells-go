package mappers

import (
	"wells-go/application/dtos"
	"wells-go/domain/entities"
)

func ToRoleResponse(role *entities.RoleEntity) *dtos.RoleResponse {
	var perms []dtos.PermissionMini
	for _, p := range role.Permissions {
		perms = append(perms, dtos.PermissionMini{
			ID:   p.ID,
			Name: p.Name,
		})
	}
	return &dtos.RoleResponse{
		ID:          role.ID,
		Name:        role.Name,
		Permissions: perms,
		CreatedAt:   role.CreatedAt,
		UpdatedAt:   role.UpdatedAt,
	}
}

func ToRoleResponses(roles []entities.RoleEntity) []*dtos.RoleResponse {
	var list []*dtos.RoleResponse
	for _, role := range roles {
		list = append(list, ToRoleResponse(&role))
	}
	return list
}
