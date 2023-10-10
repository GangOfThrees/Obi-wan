package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BusinessProfile struct {
	ID               uuid.UUID      `gorm:"column:ID;type:uuid;primaryKey;"`
	BusinessID       uuid.UUID      `gorm:"column:business_id;type:uuid;not null"`
	InitialQuestions string         `gorm:"column:initial_questions;type:json;not null"`
	CreatedAt        time.Time      `gorm:"column:CreatedAt;not null"`
	UpdatedAt        time.Time      `gorm:"column:UpdateAt;not null"`
	DeletedAt        gorm.DeletedAt `gorm:"column:DeletedAt;not null"`
}

func (BusinessProfile) TableName() string {
	return "business_profiles"
}
