package repositories

import (
	"github.com/google/uuid"
	"wells-go/domain/entities"
)

type PermissionRepository interface {
	Create(permission *entities.PermissionEntity) error
	Update(permission *entities.PermissionEntity) error
	Delete(id string) error
	FindByIDs(ids []uuid.UUID) ([]entities.PermissionEntity, error)
	FindByID(id string) (*entities.PermissionEntity, error)
	FindAll() ([]entities.PermissionEntity, error)
	FindAllWithPagination(search string, limit, offset int) ([]entities.PermissionEntity, int64, error)
}
