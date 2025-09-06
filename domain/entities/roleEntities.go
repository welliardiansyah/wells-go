package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type RoleEntity struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string         `gorm:"column:name;type:varchar(100);unique;not null" json:"name"`
	Users     []UserEntity   `gorm:"foreignKey:RoleId;references:ID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
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
