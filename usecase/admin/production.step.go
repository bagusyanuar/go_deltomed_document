package admin

import (
	"github.com/bagusyanuar/go_deltomed_document/http/request"
	"github.com/bagusyanuar/go_deltomed_document/model"
)

type ProductionStepRepository interface {
	FindAll(param string, limit int, offset int) (data []model.ProductionStep, err error)
	FindByID(id string) (data *model.ProductionStep, err error)
	Create(entity model.ProductionStep) (data *model.ProductionStep, err error)
	Patch(id string, entity model.ProductionStep) (data *model.ProductionStep, err error)
	Delete(id string) (err error)
}

type ProductionStepService interface {
	FindAll(param string, limit int, offset int) (data []model.ProductionStep, err error)
	FindByID(id string) (data *model.ProductionStep, err error)
	Create(request request.CreateProductionStepRequest) (data *model.ProductionStep, err error)
	Patch(id string, request request.CreateProductionStepRequest) (data *model.ProductionStep, err error)
	Delete(id string) (err error)
}
