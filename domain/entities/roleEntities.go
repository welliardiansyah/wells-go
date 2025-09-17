package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type RoleEntity struct {
	ID          uuid.UUID          `gorm:"type:uuid;primaryKey" json:"id"`
	Name        string             `gorm:"column:name;type:varchar(100);uniqueIndex:idx_role_name;not null" json:"name"`
	Users       []UserEntity       `gorm:"foreignKey:RoleId;references:ID"`
	Permissions []PermissionEntity `gorm:"many2many:wells_role_permissions;" json:"permissions"`
	CreatedAt   time.Time          `gorm:"index:idx_role_created_at" json:"created_at"`
	UpdatedAt   time.Time          `gorm:"index:idx_role_updated_at" json:"updated_at"`
	DeletedAt   gorm.DeletedAt     `gorm:"index:idx_role_deleted_at" json:"deleted_at,omitempty"`
}

func (RoleEntity) TableName() string {
	return "wells_roles"
}

func (r *RoleEntity) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ID == uuid.Nil {
		r.ID = uuid.New()
	}
	return
}
