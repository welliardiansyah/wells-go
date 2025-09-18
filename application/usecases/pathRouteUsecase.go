package usecases

import (
	"wells-go/domain/entities"
	"wells-go/domain/repositories"
)

type PathRouteUsecase struct {
	repo repositories.PathRouteRepository
}

func NewPathRouteUsecase(repo repositories.PathRouteRepository) PathRouteUsecase {
	return PathRouteUsecase{
		repo: repo,
	}
}

func (u *PathRouteUsecase) GetAllRoutes(filterName string, limit, offset int) ([]entities.PathRouteEntities, int64, error) {
	return u.repo.GetAllRoutes(filterName, limit, offset)
}
