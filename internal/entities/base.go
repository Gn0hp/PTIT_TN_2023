package entities

import (
	"gorm.io/gorm"
	"time"
)

type DefaultModel struct {
	ID        uint64         `json:"id" gorm:"type:INT(11) AUTO_INCREMENT;primarykey"`
	CreatedAt time.Time      `json:"created_at" gorm:"type:TIMESTAMP DEFAULT CURRENT_TIMESTAMP;index"`
	UpdatedAt time.Time      `json:"updated_at,omitempty" gorm:"type:TIMESTAMP NULL"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"type:TIMESTAMP NULL DEFAULT NULL" swaggerignore:"true"`
}

type IEntity interface {
	ToEntity() (interface{}, error)
	IsValid() bool
}
