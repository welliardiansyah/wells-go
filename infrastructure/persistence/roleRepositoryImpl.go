package persistence

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"wells-go/domain/entities"
	"wells-go/domain/repositories"
)

type RoleRepositoryImpl struct {
	db *gorm.DB
}

func NewRoleRepositoryImpl(db *gorm.DB) repositories.RoleRepository {
	return &RoleRepositoryImpl{db: db}
}

func (r *RoleRepositoryImpl) Create(role *entities.RoleEntity) error {
	return r.db.Create(role).Error
}

func (r *RoleRepositoryImpl) FindAll() ([]entities.RoleEntity, error) {
	var roles []entities.RoleEntity
	err := r.db.Preload("Permissions").Find(&roles).Error
	return roles, err
}

func (r *RoleRepositoryImpl) FindByID(id uuid.UUID) (*entities.RoleEntity, error) {
	var role entities.RoleEntity
	err := r.db.Preload("Permissions").First(&role, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *RoleRepositoryImpl) Update(role *entities.RoleEntity) error {
	return r.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(role).Error
}

func (r *RoleRepositoryImpl) Delete(id uuid.UUID) error {
	return r.db.Delete(&entities.RoleEntity{}, "id = ?", id).Error
}

func (r *RoleRepositoryImpl) FindAllWithPagination(search string, limit, offset int) ([]entities.RoleEntity, int64, error) {
	var roles []entities.RoleEntity
	var total int64

	query := r.db.Model(&entities.RoleEntity{}).Preload("Permissions")
	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if limit > 0 {
		query = query.Limit(limit).Offset(offset)
	}

	if err := query.Order("created_at desc").Find(&roles).Error; err != nil {
		return nil, 0, err
	}

	return roles, total, nil
}
