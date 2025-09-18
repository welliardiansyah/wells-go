package usecases

import (
	"github.com/google/uuid"
	"wells-go/domain/entities"
	"wells-go/domain/repositories"
)

type RouteAccessUsecase struct {
	repo repositories.RouteAccessRepository
}

func NewRouteAccessUsecase(repo repositories.RouteAccessRepository) *RouteAccessUsecase {
	return &RouteAccessUsecase{repo: repo}
}

func (u *RouteAccessUsecase) GetAll() ([]entities.RouteAccessEntities, error) {
	return u.repo.GetAll()
}

func (u *RouteAccessUsecase) GetByID(id uuid.UUID) (entities.RouteAccessEntities, error) {
	return u.repo.GetByID(id)
}

func (u *RouteAccessUsecase) Create(routeAccess *entities.RouteAccessEntities) error {
	return u.repo.Create(routeAccess)
}

func (u *RouteAccessUsecase) Update(routeAccess *entities.RouteAccessEntities) error {
	return u.repo.Update(routeAccess)
}

func (u *RouteAccessUsecase) Delete(id uuid.UUID) error {
	return u.repo.Delete(id)
}
