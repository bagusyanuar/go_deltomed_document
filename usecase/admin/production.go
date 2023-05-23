package admin

import (
	"github.com/bagusyanuar/go_deltomed_document/http/request"
	"github.com/bagusyanuar/go_deltomed_document/model"
)

type ProductionRepository interface {
	FindAll(param string, limit int, offset int) (data []model.Production, err error)
	FindByID(id string) (data *model.Production, err error)
	Create(entity model.Production) (data *model.Production, err error)
	Patch(id string, entity model.Production) (data *model.Production, err error)
	Delete(id string) (err error)
}

type ProductionService interface {
	FindAll(param string, limit int, offset int) (data []model.Production, err error)
	FindByID(id string) (data *model.Production, err error)
	Create(request request.CreateProductionRequest) (data *model.Production, err error)
	Patch(id string, request request.CreateProductionRequest) (data *model.Production, err error)
	Delete(id string) (err error)
}
