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
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
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
