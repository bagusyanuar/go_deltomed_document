package admin

import (
	"github.com/bagusyanuar/go_deltomed_document/model"
	usecase "github.com/bagusyanuar/go_deltomed_document/usecase/admin"
	"gorm.io/gorm"
)

type implementsProductionSubStepRepository struct {
	Database *gorm.DB
}

// Create implements admin.ProductionSubStepRepository
func (repository *implementsProductionSubStepRepository) Create(entity model.ProductionSubStep) (data *model.ProductionSubStep, err error) {
	panic("unimplemented")
}

// Delete implements admin.ProductionSubStepRepository
func (repository *implementsProductionSubStepRepository) Delete(id string) (err error) {
	panic("unimplemented")
}

// FindAll implements admin.ProductionSubStepRepository
func (repository *implementsProductionSubStepRepository) FindAll(param string, limit int, offset int) (data []model.ProductionSubStep, err error) {
	panic("unimplemented")
}

// FindByID implements admin.ProductionSubStepRepository
func (repository *implementsProductionSubStepRepository) FindByID(id string) (data *model.ProductionSubStep, err error) {
	panic("unimplemented")
}

// Patch implements admin.ProductionSubStepRepository
func (repository *implementsProductionSubStepRepository) Patch(id string, entity model.ProductionSubStep) (data *model.ProductionSubStep, err error) {
	panic("unimplemented")
}

func NewProductionSubStepRepository(database *gorm.DB) usecase.ProductionSubStepRepository {
	return &implementsProductionSubStepRepository{
		Database: database,
	}
}
