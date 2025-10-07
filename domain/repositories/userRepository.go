package repositories

import (
	"github.com/google/uuid"
	"wells-go/domain/entities"
)

type UserRepository interface {
	Create(user *entities.UserEntity) error
	FindByEmail(email string) (*entities.UserEntity, error)
	FindByID(id uuid.UUID) (*entities.UserEntity, error)
	Update(user *entities.UserEntity) error
	Delete(id uuid.UUID) error
	List() ([]entities.UserEntity, error)
	ListWithPagination(search string, limit, offset int) ([]entities.UserEntity, int64, error)
}
