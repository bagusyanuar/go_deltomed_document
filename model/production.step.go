package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductionStep struct {
	ID                uuid.UUID           `json:"id"`
	ProductionID      uuid.UUID           `json:"production_id"`
	Name              string              `json:"name"`
	IndexOf           uint                `json:"index_of"`
	CreatedAt         time.Time           `json:"created_at"`
	UpdatedAt         time.Time           `json:"updated_at"`
	DeletedAt         gorm.DeletedAt      `json:"deleted_at"`
	Production        Production          `gorm:"foreignKey:ProductionID" json:"production"`
	ProductionSubStep []ProductionSubStep `gorm:"foreignKey:ProductionStepID" json:"production_sub_step"`
}

func (productionStep *ProductionStep) BeforeCreate(tx *gorm.DB) (err error) {
	productionStep.ID = uuid.New()
	productionStep.CreatedAt = time.Now()
	productionStep.UpdatedAt = time.Now()
	return
}

func (ProductionStep) TableName() string {
	return "production_steps"
}
