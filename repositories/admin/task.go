package admin

import (
	"github.com/bagusyanuar/go_deltomed_document/model"
	usecase "github.com/bagusyanuar/go_deltomed_document/usecase/admin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type implementsTaskReposioty struct {
	Database *gorm.DB
}

// Create implements admin.TaskRepository
func (repository *implementsTaskReposioty) Create(entity model.Task) (data *model.Task, err error) {
	if err = repository.Database.Debug().Omit(clause.Associations).Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// Delete implements admin.TaskRepository
func (repository *implementsTaskReposioty) Delete(id string) (err error) {
	if err = repository.Database.Debug().Where("id = ?", id).Delete(&model.Task{}).Error; err != nil {
		return err
	}
	return nil
}

// FindAll implements admin.TaskRepository
func (repository *implementsTaskReposioty) FindAll(param string, limit int, offset int) (data []model.Task, err error) {
	if err = repository.Database.Debug().
		Limit(limit).
		Offset(offset).
		Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// FindByID implements admin.TaskRepository
func (repository *implementsTaskReposioty) FindByID(id string) (data *model.Task, err error) {
	if err = repository.Database.Debug().
		Where("id = ?", id).
		First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

// Patch implements admin.TaskRepository
func (repository *implementsTaskReposioty) Patch(id string, entity model.Task) (data *model.Task, err error) {
	if err = repository.Database.Debug().Omit(clause.Associations).Where("id = ?", id).Updates(entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func NewTaskRepository(database *gorm.DB) usecase.TaskRepository {
	return &implementsTaskReposioty{
		Database: database,
	}
}
