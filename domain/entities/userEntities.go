package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type UserEntity struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	Fullname  string         `gorm:"type:varchar(255);index:idx_user_fullname;not null" json:"fullname"`
	Phone     string         `gorm:"column:phone;type:varchar(15);not null" json:"phone"`
	Email     string         `gorm:"column:email;type:varchar(100);uniqueIndex:idx_user_email;not null" json:"email"`
	Password  string         `gorm:"column:password;type:varchar(100);not null" json:"password"`
	RoleId    uuid.UUID      `gorm:"type:uuid;not null" json:"role_id"`
	Role      RoleEntity     `gorm:"foreignKey:RoleId;references:ID" json:"role"`
	CreatedAt time.Time      `gorm:"index:idx_user_created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"index:idx_user_updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index:idx_user_deleted_at" json:"deleted_at,omitempty"`
}

func (UserEntity) TableName() string {
	return "wells_users"
}

func (u *UserEntity) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return
}
