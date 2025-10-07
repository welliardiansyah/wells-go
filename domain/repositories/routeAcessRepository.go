package repositories

import (
	"github.com/google/uuid"
	"wells-go/domain/entities"
)

type RouteAccessRepository interface {
	GetAll() ([]entities.RouteAccessEntities, error)
	GetByID(id uuid.UUID) (entities.RouteAccessEntities, error)
	GetAccessByRoute(method, path string) ([]entities.RouteAccessEntities, error)
	Create(routeAccess *entities.RouteAccessEntities) error
	Update(routeAccess *entities.RouteAccessEntities) error
	Delete(id uuid.UUID) error
	GetAllByRole(role string) ([]entities.RouteAccessEntities, error)
	GetAllByRoleName(roleName string) ([]entities.RouteAccessEntities, error)
	FindAllWithPagination(search string, limit, offset int) ([]entities.RouteAccessEntities, int64, error)
}
