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
}
