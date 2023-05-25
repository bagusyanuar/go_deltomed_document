package admin

import (
	"github.com/bagusyanuar/go_deltomed_document/model"
	usecase "github.com/bagusyanuar/go_deltomed_document/usecase/admin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type implementsTaskProcessRepository struct {
	Database *gorm.DB
}

// Create implements admin.TaskProcessRepository
func (repository *implementsTaskProcessRepository) Create(entity model.TaskProcess) (data *model.TaskProcess, err error) {
	if err = repository.Database.Debug().Omit(clause.Associations).Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// Delete implements admin.TaskProcessRepository
func (repository *implementsTaskProcessRepository) Delete(id string) (err error) {
	if err = repository.Database.Debug().Where("id = ?", id).Delete(&model.TaskProcess{}).Error; err != nil {
		return err
	}
	return nil
}

// FindAll implements admin.TaskProcessRepository
func (repository *implementsTaskProcessRepository) FindAll(param string, limit int, offset int) (data []model.TaskProcess, err error) {
	if err = repository.Database.Debug().
		Where("username LIKE ?", "%"+param+"%").
		Limit(limit).
		Offset(offset).
		Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// FindByID implements admin.TaskProcessRepository
func (repository *implementsTaskProcessRepository) FindByID(id string) (data *model.TaskProcess, err error) {
	if err = repository.Database.Debug().
		Where("id = ?", id).
		First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

// Patch implements admin.TaskProcessRepository
func (repository *implementsTaskProcessRepository) Patch(id string, entity model.TaskProcess) (data *model.TaskProcess, err error) {
	if err = repository.Database.Debug().Omit(clause.Associations).Where("id = ?", id).Updates(entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func NewTaskProcessRepository(database *gorm.DB) usecase.TaskProcessRepository {
	return &implementsTaskProcessRepository{
		Database: database,
	}
}
