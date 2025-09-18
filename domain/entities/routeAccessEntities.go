package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type RouteAccessEntities struct {
	ID             uuid.UUID      `gorm:"type:uuid;primaryKey" json:"ID"`
	RoutePath      string         `gorm:"column:route_path" json:"RoutePath"`
	HTTPMethod     string         `gorm:"column:http_method" json:"HTTPMethod"`
	RoleName       string         `gorm:"column:role_name" json:"RoleName"`
	PermissionName string         `gorm:"column:permission_name" json:"PermissionName"`
	CreatedAt      time.Time      `gorm:"autoCreateTime" json:"CreatedAt"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime" json:"UpdatedAt"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"DeletedAt,omitempty"`
}

func (RouteAccessEntities) TableName() string {
	return "route_access"
}

func (r *RouteAccessEntities) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ID == uuid.Nil {
		r.ID = uuid.New()
	}
	return
}
