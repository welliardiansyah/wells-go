package repositories

import "wells-go/domain/entities"

type PathRouteRepository interface {
	SeedRoute(route *entities.PathRouteEntities) error
	GetAllRoutes(filterName string, limit, offset int) ([]entities.PathRouteEntities, int64, error)
}
