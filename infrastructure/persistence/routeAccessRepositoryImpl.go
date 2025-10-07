package persistence

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strings"
	"wells-go/domain/entities"
	"wells-go/domain/repositories"
)

type RouteAccessRepositoryImpl struct {
	db *gorm.DB
}

func NewRouteAccessRepositoryImpl(db *gorm.DB) repositories.RouteAccessRepository {
	return &RouteAccessRepositoryImpl{
		db: db,
	}
}

func (r *RouteAccessRepositoryImpl) GetAll() ([]entities.RouteAccessEntities, error) {
	var routes []entities.RouteAccessEntities
	err := r.db.Find(&routes).Error
	return routes, err
}

func (r *RouteAccessRepositoryImpl) GetByID(id uuid.UUID) (entities.RouteAccessEntities, error) {
	var route entities.RouteAccessEntities
	err := r.db.First(&route, "id = ?", id).Error
	return route, err
}

func (r *RouteAccessRepositoryImpl) GetAccessByRoute(method, path string) ([]entities.RouteAccessEntities, error) {
	var access []entities.RouteAccessEntities
	err := r.db.Where("http_method = ? AND route_path = ?", method, path).Find(&access).Error
	return access, err
}

func (r *RouteAccessRepositoryImpl) Create(routeAccess *entities.RouteAccessEntities) error {
	if routeAccess.ID == uuid.Nil {
		routeAccess.ID = uuid.New()
	}
	fmt.Printf("Create RouteAccess: %+v\n", routeAccess)
	return r.db.Create(routeAccess).Error
}

func (r *RouteAccessRepositoryImpl) Update(routeAccess *entities.RouteAccessEntities) error {
	return r.db.Save(routeAccess).Error
}

func (r *RouteAccessRepositoryImpl) Delete(id uuid.UUID) error {
	var route entities.RouteAccessEntities
	err := r.db.First(&route, "id = ?", id).Error
	if err != nil {
		return err
	}
	return r.db.Delete(&route).Error
}

func (r *RouteAccessRepositoryImpl) GetAllByRole(role string) ([]entities.RouteAccessEntities, error) {
	var routes []entities.RouteAccessEntities
	err := r.db.Where("role_name = ?", role).Find(&routes).Error
	return routes, err
}

func (r *RouteAccessRepositoryImpl) GetAllByRoleName(roleName string) ([]entities.RouteAccessEntities, error) {
	var routes []entities.RouteAccessEntities
	err := r.db.Where("LOWER(role_name) = LOWER(?)", roleName).Find(&routes).Error
	return routes, err
}

func (r *RouteAccessRepositoryImpl) FindAllWithPagination(search string, limit, offset int) ([]entities.RouteAccessEntities, int64, error) {
	var routes []entities.RouteAccessEntities
	var total int64

	query := r.db.Model(&entities.RouteAccessEntities{})
	if search != "" {
		query = query.Where("LOWER(route_name) LIKE ?", "%"+strings.ToLower(search)+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if limit > 0 {
		query = query.Limit(limit).Offset(offset)
	}

	if err := query.Order("created_at desc").Find(&routes).Error; err != nil {
		return nil, 0, err
	}

	return routes, total, nil
}
