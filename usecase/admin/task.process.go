package admin

import (
	"github.com/bagusyanuar/go_deltomed_document/http/request"
	"github.com/bagusyanuar/go_deltomed_document/model"
)

type TaskProcessRepository interface {
	FindAll(param string, limit int, offset int) (data []model.TaskProcess, err error)
	FindByID(id string) (data *model.TaskProcess, err error)
	Create(entity model.TaskProcess) (data *model.TaskProcess, err error)
	Patch(id string, entity model.TaskProcess) (data *model.TaskProcess, err error)
	Delete(id string) (err error)
}

type TaskProcessService interface {
	FindAll(param string, limit int, offset int) (data []model.TaskProcess, err error)
	FindByID(id string) (data *model.TaskProcess, err error)
	Create(request request.CreateTaskProcessRequest) (data *model.TaskProcess, err error)
	Patch(id string, request request.CreateTaskProcessRequest) (data *model.TaskProcess, err error)
	Delete(id string) (err error)
}
