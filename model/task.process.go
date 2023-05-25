package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaskProcess struct {
	ID                  uuid.UUID         `json:"id"`
	TaskID              uuid.UUID         `json:"task_id"`
	ProductionStepID    uuid.UUID         `json:"production_step_id"`
	ProductionsubStepID uuid.UUID         `json:"production_sub_step_id"`
	SubmittedAt         *time.Time        `json:"submitted_at"`
	ApprovedAt          time.Time         `json:"approved_at"`
	Image               *string           `json:"image"`
	CreatedAt           time.Time         `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt           time.Time         `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt           gorm.DeletedAt    `gorm:"column:deleted_at;" json:"deleted_at"`
	Task                Task              `gorm:"foreignKey:TaskID"`
	ProductionStep      ProductionStep    `gorm:"foreignKey:ProductionStepID"`
	ProductionSubStep   ProductionSubStep `gorm:"foreignKey:ProductionSubStepID"`
}

func (taskProcess *TaskProcess) BeforeCreate(tx *gorm.DB) (err error) {
	taskProcess.ID = uuid.New()
	taskProcess.CreatedAt = time.Now()
	taskProcess.UpdatedAt = time.Now()
	return
}

func (TaskProcess) TableName() string {
	return "task_processes"
}
