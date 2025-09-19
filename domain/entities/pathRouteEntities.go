package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PathRouteEntities struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string    `gorm:"size:255;not null" json:"name"`
	Path      string    `gorm:"size:255;not null" json:"path"`
	Method    string    `gorm:"size:10;not null" json:"method"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
