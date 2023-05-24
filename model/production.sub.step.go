package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductionSubStep struct {
	ID               uuid.UUID      `json:"id"`
	ProductionStepID uuid.UUID      `json:"production_step_id"`
	Name             string         `json:"name"`
	IndexOf          uint           `json:"index_of"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `json:"deleted_at"`
	ProductionStep   ProductionStep `gorm:"foreignKey:ProductionStepID"`
}

func (productionSubStep *ProductionSubStep) BeforeCreate(tx *gorm.DB) (err error) {
	productionSubStep.ID = uuid.New()
	productionSubStep.CreatedAt = time.Now()
	productionSubStep.UpdatedAt = time.Now()
	return
}

func (ProductionSubStep) TableName() string {
	return "production_sub_steps"
}
