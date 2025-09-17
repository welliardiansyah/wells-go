package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type PermissionEntity struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string         `gorm:"column:name;type:varchar(100);uniqueIndex:idx_permission_name;not null" json:"name"`
	CanCreate bool           `gorm:"default:false" json:"can_create"`
	CanRead   bool           `gorm:"default:false" json:"can_read"`
	CanUpdate bool           `gorm:"default:false" json:"can_update"`
	CanDelete bool           `gorm:"default:false" json:"can_delete"`
	CanExport bool           `gorm:"default:false" json:"can_export"`
	CanImport bool           `gorm:"default:false" json:"can_import"`
	CanView   bool           `gorm:"default:false" json:"can_view"`
	Roles     []RoleEntity   `gorm:"many2many:wells_role_permissions;" json:"roles"`
	CreatedAt time.Time      `gorm:"index:idx_permission_created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"index:idx_permission_updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index:idx_permission_deleted_at" json:"deleted_at,omitempty"`
}

func (PermissionEntity) TableName() string {
	return "wells_permissions"
}

func (p *PermissionEntity) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return
}
