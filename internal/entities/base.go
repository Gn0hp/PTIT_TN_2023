package entities

import (
	"gorm.io/gorm"
	"time"
)

type DefaultModel struct {
	gorm.Model
	ID        uint64         `json:"id" gorm:"AUTO_INCREMENT not null;primaryKey"`
	CreatedAt time.Time      `json:"created_at" gorm:"type:TIMESTAMP DEFAULT CURRENT_TIMESTAMP;index;autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at,omitempty" gorm:"type:TIMESTAMP NULL;autoUpdateTime:milli"`
	IsDeleted bool           `json:"is_deleted,omitempty" gorm:"default:false"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"type:TIMESTAMP NULL DEFAULT NULL " swaggerignore:"true"`
}

type IEntity interface {
	ToEntity() (interface{}, error)
	IsValid() bool
}
