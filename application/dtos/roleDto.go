package dtos

import (
	"github.com/google/uuid"
	"time"
)

type CreateRoleRequest struct {
	Name          string      `json:"name" binding:"required"`
	PermissionIDs []uuid.UUID `json:"permission_ids" binding:"required"`
}

type UpdateRoleRequest struct {
	Name          string      `json:"name" binding:"required"`
	PermissionIDs []uuid.UUID `json:"permission_ids" binding:"required"`
}

type RoleResponse struct {
	ID          uuid.UUID        `json:"id"`
	Name        string           `json:"name"`
	Permissions []PermissionMini `json:"permissions"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
}

type PermissionMini struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
