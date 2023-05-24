package admin

import (
	"github.com/bagusyanuar/go_deltomed_document/model"
	usecase "github.com/bagusyanuar/go_deltomed_document/usecase/admin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type implementsProductionSubStepRepository struct {
	Database *gorm.DB
}

// Create implements admin.ProductionSubStepRepository
func (repository *implementsProductionSubStepRepository) Create(entity model.ProductionSubStep) (data *model.ProductionSubStep, err error) {
	if err = repository.Database.Debug().Omit(clause.Associations).Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// Delete implements admin.ProductionSubStepRepository
func (repository *implementsProductionSubStepRepository) Delete(id string) (err error) {
	if err = repository.Database.Debug().Where("id = ?", id).Delete(&model.ProductionSubStep{}).Error; err != nil {
		return err
	}
	return nil
}

// FindAll implements admin.ProductionSubStepRepository
func (repository *implementsProductionSubStepRepository) FindAll(param string, limit int, offset int) (data []model.ProductionSubStep, err error) {
	if err = repository.Database.Debug().
		Where("username LIKE ?", "%"+param+"%").
		Limit(limit).
		Offset(offset).
		Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// FindByID implements admin.ProductionSubStepRepository
func (repository *implementsProductionSubStepRepository) FindByID(id string) (data *model.ProductionSubStep, err error) {
	if err = repository.Database.Debug().
		Where("id = ?", id).
		First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

// Patch implements admin.ProductionSubStepRepository
func (repository *implementsProductionSubStepRepository) Patch(id string, entity model.ProductionSubStep) (data *model.ProductionSubStep, err error) {
	if err = repository.Database.Debug().Omit(clause.Associations).Where("id = ?", id).Updates(entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func NewProductionSubStepRepository(database *gorm.DB) usecase.ProductionSubStepRepository {
	return &implementsProductionSubStepRepository{
		Database: database,
	}
}
