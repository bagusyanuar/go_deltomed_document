package admin

import (
	"github.com/bagusyanuar/go_deltomed_document/model"
	usecase "github.com/bagusyanuar/go_deltomed_document/usecase/admin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type implementsProductionRepository struct {
	Database *gorm.DB
}

// Create implements admin.ProductionRepository
func (repository *implementsProductionRepository) Create(entity model.Production) (data *model.Production, err error) {
	if err = repository.Database.Debug().Omit(clause.Associations).Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// Delete implements admin.ProductionRepository
func (repository *implementsProductionRepository) Delete(id string) (err error) {
	if err = repository.Database.Debug().Where("id = ?", id).Delete(&model.Production{}).Error; err != nil {
		return err
	}
	return nil
}

// FindAll implements admin.ProductionRepository
func (repository *implementsProductionRepository) FindAll(param string, limit int, offset int) (data []model.Production, err error) {
	if err = repository.Database.Debug().
		Where("name LIKE ?", "%"+param+"%").
		Limit(limit).
		Offset(offset).
		Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// FindByID implements admin.ProductionRepository
func (repository *implementsProductionRepository) FindByID(id string) (data *model.Production, err error) {
	if err = repository.Database.Debug().
		Where("id = ?", id).
		First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

// Patch implements admin.ProductionRepository
func (repository *implementsProductionRepository) Patch(id string, entity model.Production) (data *model.Production, err error) {
	if err = repository.Database.Debug().Omit(clause.Associations).Where("id = ?", id).Updates(entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func NewProductionRepository(database *gorm.DB) usecase.ProductionRepository {
	return &implementsProductionRepository{
		Database: database,
	}
}
