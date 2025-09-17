package persistence

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"wells-go/domain/entities"
)

type PermissionRepositoryImpl struct {
	db *gorm.DB
}

func NewPermissionRepositoryImpl(db *gorm.DB) *PermissionRepositoryImpl {
	return &PermissionRepositoryImpl{
		db: db,
	}
}

func (r *PermissionRepositoryImpl) Create(permission *entities.PermissionEntity) error {
	return r.db.Create(permission).Error
}

func (r *PermissionRepositoryImpl) Update(permission *entities.PermissionEntity) error {
	return r.db.Save(permission).Error
}

func (r *PermissionRepositoryImpl) Delete(id string) error {
	return r.db.Delete(&entities.PermissionEntity{}, "id = ?", id).Error
}

func (r *PermissionRepositoryImpl) FindByID(id string) (*entities.PermissionEntity, error) {
	var permission entities.PermissionEntity
	if err := r.db.First(&permission, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &permission, nil
}

func (r *PermissionRepositoryImpl) FindAll() ([]entities.PermissionEntity, error) {
	var permissions []entities.PermissionEntity
	if err := r.db.Find(&permissions).Error; err != nil {
		return nil, err
	}
	return permissions, nil
}

func (r *PermissionRepositoryImpl) FindByIDs(ids []uuid.UUID) ([]entities.PermissionEntity, error) {
	var permissions []entities.PermissionEntity
	err := r.db.Where("id IN ?", ids).Find(&permissions).Error
	return permissions, err
}
