package admin

import (
	"github.com/bagusyanuar/go_deltomed_document/common"
	"github.com/bagusyanuar/go_deltomed_document/http/request"
	"github.com/bagusyanuar/go_deltomed_document/model"
	usecase "github.com/bagusyanuar/go_deltomed_document/usecase/admin"
	"github.com/google/uuid"
)

type implementsProductionSubStepService struct {
	ProductionSubStepRepository usecase.ProductionSubStepRepository
}

// Create implements admin.ProductionSubStepService
func (service *implementsProductionSubStepService) Create(request request.CreateProductionSubStepRequest) (data *model.ProductionSubStep, err error) {
	productionStepID, err := uuid.Parse(request.ProductionStepID)
	if err != nil {
		return nil, err
	}
	entity := model.ProductionSubStep{
		Name:             request.Name,
		ProductionStepID: productionStepID,
		IndexOf:          request.IndexOf,
	}
	return service.ProductionSubStepRepository.Create(entity)
}

// Delete implements admin.ProductionSubStepService
func (service *implementsProductionSubStepService) Delete(id string) (err error) {
	return service.ProductionSubStepRepository.Delete(id)
}

// FindAll implements admin.ProductionSubStepService
func (service *implementsProductionSubStepService) FindAll(param string, limit int, offset int) (data []model.ProductionSubStep, err error) {
	if limit == 0 {
		limit = common.DefaultLimit
	}
	return service.ProductionSubStepRepository.FindAll(param, limit, offset)
}

// FindByID implements admin.ProductionSubStepService
func (service *implementsProductionSubStepService) FindByID(id string) (data *model.ProductionSubStep, err error) {
	return service.ProductionSubStepRepository.FindByID(id)
}

// Patch implements admin.ProductionSubStepService
func (service *implementsProductionSubStepService) Patch(id string, request request.CreateProductionSubStepRequest) (data *model.ProductionSubStep, err error) {
	productionStepID, err := uuid.Parse(request.ProductionStepID)
	if err != nil {
		return nil, err
	}
	entity := model.ProductionSubStep{
		Name:             request.Name,
		ProductionStepID: productionStepID,
		IndexOf:          request.IndexOf,
	}
	return service.ProductionSubStepRepository.Patch(id, entity)
}

func NewProductionSubStepService(productionSubStepRepository usecase.ProductionSubStepRepository) usecase.ProductionSubStepService {
	return &implementsProductionSubStepService{
		ProductionSubStepRepository: productionSubStepRepository,
	}
}
