package usecases

import (
	"github.com/google/uuid"
	"wells-go/application/dtos"
	"wells-go/application/mappers"
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

func (u *RouteAccessUsecase) GetAllByRole(role string) ([]entities.RouteAccessEntities, error) {
	return u.repo.GetAllByRole(role)
}

func (u *RouteAccessUsecase) GetAllByRoleName(roleName string) ([]entities.RouteAccessEntities, error) {
	return u.repo.GetAllByRoleName(roleName)
}

func (u *RouteAccessUsecase) GetAllWithPagination(search string, limit, offset int) ([]dtos.RouteAccessResponse, int64, error) {
	routes, total, err := u.repo.FindAllWithPagination(search, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	res := make([]dtos.RouteAccessResponse, len(routes))
	for i, r := range routes {
		res[i] = mappers.ToRouteAccessResponse(&r)
	}

	return res, total, nil
}
