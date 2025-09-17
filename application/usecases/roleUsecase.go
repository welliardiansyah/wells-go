package usecases

import (
	"fmt"
	"github.com/google/uuid"
	"wells-go/application/dtos"
	"wells-go/application/mappers"
	"wells-go/domain/entities"
	"wells-go/domain/repositories"
)

type RoleUsecase struct {
	repo           repositories.RoleRepository
	repoPermission repositories.PermissionRepository
}

func NewRoleUsecase(repo repositories.RoleRepository, repoPermission repositories.PermissionRepository) *RoleUsecase {
	return &RoleUsecase{repo: repo, repoPermission: repoPermission}
}

func (u *RoleUsecase) CreateRole(req dtos.CreateRoleRequest) (*dtos.RoleResponse, error) {
	if len(req.PermissionIDs) == 0 {
		return nil, fmt.Errorf("role must have at least one permission")
	}

	role := &entities.RoleEntity{
		Name: req.Name,
	}

	permissions, err := u.repoPermission.FindByIDs(req.PermissionIDs)
	if err != nil {
		return nil, err
	}
	if len(permissions) == 0 {
		return nil, fmt.Errorf("no valid permissions found for given IDs")
	}

	role.Permissions = permissions

	if err := u.repo.Create(role); err != nil {
		return nil, err
	}

	return mappers.ToRoleResponse(role), nil
}

func (u *RoleUsecase) GetAllRoles() ([]*dtos.RoleResponse, error) {
	roles, err := u.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return mappers.ToRoleResponses(roles), nil
}

func (u *RoleUsecase) GetRoleByID(id uuid.UUID) (*dtos.RoleResponse, error) {
	role, err := u.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return mappers.ToRoleResponse(role), nil
}

func (u *RoleUsecase) UpdateRole(id uuid.UUID, req dtos.UpdateRoleRequest) (*dtos.RoleResponse, error) {
	role, err := u.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	role.Name = req.Name

	var permissions []entities.PermissionEntity
	for _, pid := range req.PermissionIDs {
		permissions = append(permissions, entities.PermissionEntity{ID: pid})
	}
	role.Permissions = permissions

	if err := u.repo.Update(role); err != nil {
		return nil, err
	}
	return mappers.ToRoleResponse(role), nil
}

func (u *RoleUsecase) DeleteRole(id uuid.UUID) error {
	return u.repo.Delete(id)
}
