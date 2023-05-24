package admin

import (
	"github.com/bagusyanuar/go_deltomed_document/http/request"
	"github.com/bagusyanuar/go_deltomed_document/model"
)

type ProductionSubStepRepository interface {
	FindAll(param string, limit int, offset int) (data []model.ProductionSubStep, err error)
	FindByID(id string) (data *model.ProductionSubStep, err error)
	Create(entity model.ProductionSubStep) (data *model.ProductionSubStep, err error)
	Patch(id string, entity model.ProductionSubStep) (data *model.ProductionSubStep, err error)
	Delete(id string) (err error)
}

type ProductionSubStepService interface {
	FindAll(param string, limit int, offset int) (data []model.ProductionSubStep, err error)
	FindByID(id string) (data *model.ProductionSubStep, err error)
	Create(request request.CreateProductionSubStepRequest) (data *model.ProductionSubStep, err error)
	Patch(id string, request request.CreateProductionSubStepRequest) (data *model.ProductionSubStep, err error)
	Delete(id string) (err error)
}
