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
