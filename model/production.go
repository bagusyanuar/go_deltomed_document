package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Production struct {
	ID             uuid.UUID        `json:"id"`
	Code           string           `json:"code"`
	Name           string           `json:"name"`
	Date           datatypes.Date   `json:"date"`
	CreatedAt      time.Time        `json:"created_at"`
	UpdatedAt      time.Time        `json:"updated_at"`
	DeletedAt      gorm.DeletedAt   `json:"deleted_at"`
	ProductionStep []ProductionStep `gorm:"foreignKey:ProductionID" json:"production_step"`
}

func (production *Production) BeforeCreate(tx *gorm.DB) (err error) {
	production.ID = uuid.New()
	production.CreatedAt = time.Now()
	production.UpdatedAt = time.Now()
	return
}

func (Production) TableName() string {
	return "productions"
}
