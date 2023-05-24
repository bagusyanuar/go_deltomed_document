package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Task struct {
	ID           uuid.UUID      `json:"id"`
	ProductionID uuid.UUID      `json:"production_id"`
	SupportID    uuid.UUID      `json:"support_id"`
	Code         string         `json:"code"`
	Name         string         `json:"name"`
	StartDate    datatypes.Date `json:"start_date"`
	FinishDate   datatypes.Date `json:"finish_date"`
	CompletedAt  datatypes.Date `json:"completed_at"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
	Production   Production     `gorm:"foreignKey:ProductionID"`
	Support      User           `gorm:"foreignKey:SupportID"`
}

func (task *Task) BeforeCreate(tx *gorm.DB) (err error) {
	task.ID = uuid.New()
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	return
}

func (Task) TableName() string {
	return "tasks"
}
