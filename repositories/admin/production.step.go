package admin

import (
	"github.com/bagusyanuar/go_deltomed_document/model"
	usecase "github.com/bagusyanuar/go_deltomed_document/usecase/admin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type implementsProductionStepRepository struct {
	Database *gorm.DB
}

// Create implements admin.ProductionStepRepository
func (repository *implementsProductionStepRepository) Create(entity model.ProductionStep) (data *model.ProductionStep, err error) {
	if err = repository.Database.Debug().Omit(clause.Associations).Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// Delete implements admin.ProductionStepRepository
func (repository *implementsProductionStepRepository) Delete(id string) (err error) {
	if err = repository.Database.Debug().Where("id = ?", id).Delete(&model.ProductionStep{}).Error; err != nil {
		return err
	}
	return nil
}

// FindAll implements admin.ProductionStepRepository
func (repository *implementsProductionStepRepository) FindAll(param string, limit int, offset int) (data []model.ProductionStep, err error) {
	if err = repository.Database.Debug().
		Where("username LIKE ?", "%"+param+"%").
		Limit(limit).
		Offset(offset).
		Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// FindByID implements admin.ProductionStepRepository
func (repository *implementsProductionStepRepository) FindByID(id string) (data *model.ProductionStep, err error) {
	if err = repository.Database.Debug().
		Where("id = ?", id).
		First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

// Patch implements admin.ProductionStepRepository
func (repository *implementsProductionStepRepository) Patch(id string, entity model.ProductionStep) (data *model.ProductionStep, err error) {
	if err = repository.Database.Debug().Omit(clause.Associations).Where("id = ?", id).Updates(entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func NewProductionStepRepository(database *gorm.DB) usecase.ProductionStepRepository {
	return &implementsProductionStepRepository{
		Database: database,
	}
}
