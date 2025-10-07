package persistence

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"wells-go/domain/entities"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) Create(user *entities.UserEntity) error {
	return r.db.Create(user).Error
}

func (r *UserRepositoryImpl) FindByEmail(email string) (*entities.UserEntity, error) {
	var user entities.UserEntity
	err := r.db.
		Preload("Role").
		Preload("Role.Permissions").
		Where("email = ?", email).
		First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) FindByID(id uuid.UUID) (*entities.UserEntity, error) {
	var user entities.UserEntity
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) Update(user *entities.UserEntity) error {
	return r.db.Save(user).Error
}

func (r *UserRepositoryImpl) Delete(id uuid.UUID) error {
	return r.db.Delete(&entities.UserEntity{}, "id = ?", id).Error
}

func (r *UserRepositoryImpl) List() ([]entities.UserEntity, error) {
	var users []entities.UserEntity
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepositoryImpl) ListWithPagination(search string, limit, offset int) ([]entities.UserEntity, int64, error) {
	var users []entities.UserEntity
	var total int64

	query := r.db.Model(&entities.UserEntity{}).Preload("Role").Preload("Role.Permissions")

	if search != "" {
		query = query.Where("fullname LIKE ? OR email LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if limit > 0 {
		query = query.Limit(limit).Offset(offset)
	}

	if err := query.Order("created_at desc").Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}
