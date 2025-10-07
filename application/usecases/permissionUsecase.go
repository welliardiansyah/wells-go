package usecases

import (
	"wells-go/application/dtos"
	"wells-go/application/mappers"
	"wells-go/domain/entities"
	"wells-go/domain/repositories"
)

type PermissionUsecase struct {
	repo repositories.PermissionRepository
}

func NewPermissionUsecase(repo repositories.PermissionRepository) *PermissionUsecase {
	return &PermissionUsecase{repo: repo}
}

func (u *PermissionUsecase) Create(req dtos.CreatePermissionRequest) (dtos.PermissionResponse, error) {
	permission := entities.PermissionEntity{
		Name:      req.Name,
		CanCreate: req.CanCreate,
		CanRead:   req.CanRead,
		CanUpdate: req.CanUpdate,
		CanDelete: req.CanDelete,
		CanExport: req.CanExport,
		CanImport: req.CanImport,
		CanView:   req.CanView,
	}
	if err := u.repo.Create(&permission); err != nil {
		return dtos.PermissionResponse{}, err
	}
	return mappers.ToPermissionResponse(permission), nil
}

func (u *PermissionUsecase) Update(id string, req dtos.UpdatePermissionRequest) (dtos.PermissionResponse, error) {
	permission, err := u.repo.FindByID(id)
	if err != nil {
		return dtos.PermissionResponse{}, err
	}
	permission.Name = req.Name
	permission.CanCreate = req.CanCreate
	permission.CanRead = req.CanRead
	permission.CanUpdate = req.CanUpdate
	permission.CanDelete = req.CanDelete
	permission.CanExport = req.CanExport
	permission.CanImport = req.CanImport
	permission.CanView = req.CanView

	if err := u.repo.Update(permission); err != nil {
		return dtos.PermissionResponse{}, err
	}
	return mappers.ToPermissionResponse(*permission), nil
}

func (u *PermissionUsecase) Delete(id string) error {
	return u.repo.Delete(id)
}

func (u *PermissionUsecase) FindByID(id string) (dtos.PermissionResponse, error) {
	permission, err := u.repo.FindByID(id)
	if err != nil {
		return dtos.PermissionResponse{}, err
	}
	return mappers.ToPermissionResponse(*permission), nil
}

func (u *PermissionUsecase) FindAll() ([]dtos.PermissionResponse, error) {
	permissions, err := u.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return mappers.ToPermissionResponseList(permissions), nil
}

func (u *PermissionUsecase) FindAllWithPagination(search string, limit, offset int) ([]dtos.PermissionResponse, int64, error) {
	permissions, total, err := u.repo.FindAllWithPagination(search, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	res := make([]dtos.PermissionResponse, len(permissions))
	for i, p := range permissions {
		res[i] = mappers.ToPermissionResponse(p)
	}

	return res, total, nil
}
