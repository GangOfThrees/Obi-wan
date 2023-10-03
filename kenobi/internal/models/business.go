package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Business struct {
	ID                  uuid.UUID      `gorm:"type:uuid;primaryKey"`
	Name                string         `gorm:"column:name"`
	Description         string         `gorm:"column:description"`
	Address             string         `gorm:"column:address"`
	City                string         `gorm:"column:city"`
	State               string         `gorm:"column:state"`
	Zip                 string         `gorm:"column:zip"`
	Phone               string         `gorm:"column:phone"`
	Website             string         `gorm:"column:website"`
	Email               string         `gorm:"column:email"`
	Status              string         `gorm:"column:status"`
	PreferredBotService string         `gorm:"column:preferred_bot_service"`
	CreatedAt           time.Time      `gorm:"column:CreatedAt"`
	UpdatedAt           time.Time      `gorm:"column:UpdatedAt"`
	DeletedAt           gorm.DeletedAt `gorm:"column:DeletedAt"`
}

func (b *Business) TableName() string {
	return "businesses"
}
