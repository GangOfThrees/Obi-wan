package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ApiKey struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	BusinessID uuid.UUID `gorm:"type:uuid;not null"`
	Key        string    `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}

func (ApiKey) TableName() string {
	return "api_keys"
}
