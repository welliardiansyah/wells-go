package repositories

import (
	"github.com/google/uuid"
	"wells-go/domain/entities"
)

type RoleRepository interface {
	Create(role *entities.RoleEntity) error
	FindAll() ([]entities.RoleEntity, error)
	FindByID(id uuid.UUID) (*entities.RoleEntity, error)
	Update(role *entities.RoleEntity) error
	Delete(id uuid.UUID) error
}
