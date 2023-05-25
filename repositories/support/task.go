package support

import (
	"github.com/bagusyanuar/go_deltomed_document/model"
	usecase "github.com/bagusyanuar/go_deltomed_document/usecase/support"
	"gorm.io/gorm"
)

type implementsTaskRepository struct {
	Database *gorm.DB
}

// GetData implements support.TaskRepository
func (repository *implementsTaskRepository) GetData(authorizedID string, param string, limit int, offset int) (data []model.Task, err error) {
	if err = repository.Database.Debug().
		Where("support_id = ?", authorizedID).
		Limit(limit).
		Offset(offset).
		Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func NewTaskRepository(database *gorm.DB) usecase.TaskRepository {
	return &implementsTaskRepository{
		Database: database,
	}
}
