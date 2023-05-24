package admin

import (
	"github.com/bagusyanuar/go_deltomed_document/common"
	"github.com/bagusyanuar/go_deltomed_document/http/request"
	"github.com/bagusyanuar/go_deltomed_document/model"
	usecase "github.com/bagusyanuar/go_deltomed_document/usecase/admin"
	"github.com/google/uuid"
)

type implementsProductionStepService struct {
	ProductionStepRepository usecase.ProductionStepRepository
}

// Create implements admin.ProductionStepService
func (service *implementsProductionStepService) Create(request request.CreateProductionStepRequest) (data *model.ProductionStep, err error) {
	productionID, err := uuid.Parse(request.ProductionID)
	if err != nil {
		return nil, err
	}
	entity := model.ProductionStep{
		Name:         request.Name,
		ProductionID: productionID,
		IndexOf:      request.IndexOf,
	}
	return service.ProductionStepRepository.Create(entity)
}

// Delete implements admin.ProductionStepService
func (service *implementsProductionStepService) Delete(id string) (err error) {
	return service.ProductionStepRepository.Delete(id)
}

// FindAll implements admin.ProductionStepService
func (service *implementsProductionStepService) FindAll(param string, limit int, offset int) (data []model.ProductionStep, err error) {
	if limit == 0 {
		limit = common.DefaultLimit
	}
	return service.ProductionStepRepository.FindAll(param, limit, offset)
}

// FindByID implements admin.ProductionStepService
func (service *implementsProductionStepService) FindByID(id string) (data *model.ProductionStep, err error) {
	return service.ProductionStepRepository.FindByID(id)
}

// Patch implements admin.ProductionStepService
func (service *implementsProductionStepService) Patch(id string, request request.CreateProductionStepRequest) (data *model.ProductionStep, err error) {
	productionID, err := uuid.Parse(request.ProductionID)
	if err != nil {
		return nil, err
	}
	entity := model.ProductionStep{
		Name:         request.Name,
		ProductionID: productionID,
		IndexOf:      request.IndexOf,
	}
	return service.ProductionStepRepository.Patch(id, entity)
}

func NewProductionStepService(productionStepRepository usecase.ProductionStepRepository) usecase.ProductionStepService {
	return &implementsProductionStepService{
		ProductionStepRepository: productionStepRepository,
	}
}
