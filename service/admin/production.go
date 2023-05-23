package admin

import (
	"time"

	"github.com/bagusyanuar/go_deltomed_document/common"
	"github.com/bagusyanuar/go_deltomed_document/http/request"
	"github.com/bagusyanuar/go_deltomed_document/model"
	usecase "github.com/bagusyanuar/go_deltomed_document/usecase/admin"
	"gorm.io/datatypes"
)

type implementsProductionService struct {
	ProductionRepository usecase.ProductionRepository
}

// Create implements admin.ProductionService
func (service *implementsProductionService) Create(request request.CreateProductionRequest) (data *model.Production, err error) {
	date, err := time.Parse(common.DateFormat, request.Date)
	if err != nil {
		return nil, err
	}
	entity := model.Production{
		Name: request.Name,
		Date: datatypes.Date(date),
	}
	return service.ProductionRepository.Create(entity)
}

// Delete implements admin.ProductionService
func (service *implementsProductionService) Delete(id string) (err error) {
	return service.ProductionRepository.Delete(id)
}

// FindAll implements admin.ProductionService
func (service *implementsProductionService) FindAll(param string, limit int, offset int) (data []model.Production, err error) {
	if limit == 0 {
		limit = common.DefaultLimit
	}
	return service.ProductionRepository.FindAll(param, limit, offset)
}

// FindByID implements admin.ProductionService
func (service *implementsProductionService) FindByID(id string) (data *model.Production, err error) {
	return service.ProductionRepository.FindByID(id)
}

// Patch implements admin.ProductionService
func (service *implementsProductionService) Patch(id string, request request.CreateProductionRequest) (data *model.Production, err error) {
	date, err := time.Parse(common.DateFormat, request.Date)
	if err != nil {
		return nil, err
	}
	entity := model.Production{
		Name: request.Name,
		Date: datatypes.Date(date),
	}
	return service.ProductionRepository.Patch(id, entity)
}

func NewProductionService(productionRepository usecase.ProductionRepository) usecase.ProductionService {
	return &implementsProductionService{
		ProductionRepository: productionRepository,
	}
}
