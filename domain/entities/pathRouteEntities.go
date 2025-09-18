package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PathRouteEntities struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name      string    `gorm:"size:255;not null"`
	Path      string    `gorm:"size:255;not null"`
	Method    string    `gorm:"size:10;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (PathRouteEntities) TableName() string {
	return "wells_path_route"
}

func (r *PathRouteEntities) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ID == uuid.Nil {
		r.ID = uuid.New()
	}
	return
}
