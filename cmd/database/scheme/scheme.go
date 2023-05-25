package scheme

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	Email     string         `gorm:"index:idx_email,unique;type:varchar(255);" json:"email"`
	Username  string         `gorm:"index:idx_username,unique;type:varchar(255);not null" json:"username"`
	Password  *string        `gorm:"type:text" json:"password"`
	Roles     datatypes.JSON `gorm:"type:longtext;not null" json:"roles"`
	CreatedAt time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;" json:"deleted_at"`
}

type Production struct {
	ID        uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	Code      string         `gorm:"index:idx_code,unique;type:varchar(16)" json:"code"`
	Name      string         `gorm:"type:varchar(255);not null;" json:"name"`
	Date      datatypes.Date `gorm:"type:date;not null;" json:"date"`
	CreatedAt time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;" json:"deleted_at"`
}

type ProductionStep struct {
	ID           uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	ProductionID uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_production_id;not null;" json:"production_id"`
	Name         string         `gorm:"type:varchar(255);not null;" json:"name"`
	IndexOf      uint           `gorm:"type:int(11);default:0" json:"index_of"`
	CreatedAt    time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;" json:"deleted_at"`
	Production   Production     `gorm:"foreignKey:ProductionID"`
}

type ProductionSubStep struct {
	ID               uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	ProductionStepID uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_production_step_id;not null;" json:"production_step_id"`
	Name             string         `gorm:"type:varchar(255);not null;" json:"name"`
	IndexOf          uint           `gorm:"type:int(11);default:0" json:"index_of"`
	CreatedAt        time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt        time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at;" json:"deleted_at"`
	ProductionStep   ProductionStep `gorm:"foreignKey:ProductionStepID"`
}

type Task struct {
	ID           uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	ProductionID uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_production_id;not null;" json:"production_id"`
	SupportID    uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_support_id;not null;" json:"support_id"`
	Code         string         `gorm:"index:idx_code,unique;type:varchar(16);not null;" json:"code"`
	Name         string         `gorm:"type:varchar(255);not null;" json:"name"`
	StartDate    datatypes.Date `gorm:"type:date;not null;" json:"start_date"`
	FinishDate   datatypes.Date `gorm:"type:date;not null;" json:"finish_date"`
	CompletedAt  datatypes.Date `gorm:"type:date;" json:"completed_at"`
	CreatedAt    time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;" json:"deleted_at"`
	Production   Production     `gorm:"foreignKey:ProductionID"`
	Support      User           `gorm:"foreignKey:SupportID"`
}

type TaskProcess struct {
	ID                  uuid.UUID         `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	TaskID              uuid.UUID         `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_task;not null;" json:"task_id"`
	ProductionStepID    uuid.UUID         `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_production_step_id;not null;" json:"production_step_id"`
	ProductionSubStepID uuid.UUID         `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_production_sub_step_id;not null;" json:"production_sub_step_id"`
	SubmittedAt         time.Time         `gorm:"column:submitted_at;not null;" json:"submitted_at"`
	ApprovedAt          time.Time         `gorm:"column:approved_at;" json:"approved_at"`
	Image               string            `gorm:"type:text;not null;" json:"image"`
	CreatedAt           time.Time         `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt           time.Time         `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt           gorm.DeletedAt    `gorm:"column:deleted_at;" json:"deleted_at"`
	Task                Task              `gorm:"foreignKey:TaskID"`
	ProductionStep      ProductionStep    `gorm:"foreignKey:ProductionStepID"`
	ProductionSubStep   ProductionSubStep `gorm:"foreignKey:ProductionSubStepID"`
}
