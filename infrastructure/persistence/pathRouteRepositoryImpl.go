package persistence

import (
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"wells-go/domain/entities"
	"wells-go/domain/repositories"
)

type PathRouteRepositoryImpl struct {
	db *gorm.DB
}

func NewPathRouteRepositoryImpl(db *gorm.DB) repositories.PathRouteRepository {
	return &PathRouteRepositoryImpl{db: db}
}

func (r *PathRouteRepositoryImpl) SeedRoute(route *entities.PathRouteEntities) error {
	var existing entities.PathRouteEntities
	err := r.db.Where("path = ? AND method = ?", route.Path, route.Method).First(&existing).Error
	if err == nil {
		log.Info().Str("path", route.Path).Str("method", route.Method).Msg("Route already exists")
		return nil
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error().Err(err).Msg("DB error checking existing route")
		return err
	}

	if err := r.db.Create(route).Error; err != nil {
		log.Error().Err(err).Str("path", route.Path).Str("method", route.Method).Msg("Failed to insert route")
		return err
	}

	log.Info().Str("path", route.Path).Str("method", route.Method).Msg("New route seeded")
	return nil
}

func (r *PathRouteRepositoryImpl) GetAllRoutes(filterName string, limit, offset int) ([]entities.PathRouteEntities, int64, error) {
	var routes []entities.PathRouteEntities
	var total int64

	query := r.db.Model(&entities.PathRouteEntities{})

	if filterName != "" {
		query = query.Where("path LIKE ?", "%"+filterName+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if limit > 0 {
		query = query.Limit(limit).Offset(offset)
	}

	if err := query.Order("id desc").Find(&routes).Error; err != nil {
		return nil, 0, err
	}

	return routes, total, nil
}
