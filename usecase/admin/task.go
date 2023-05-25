package admin

import (
	"github.com/bagusyanuar/go_deltomed_document/http/request"
	"github.com/bagusyanuar/go_deltomed_document/model"
)

type TaskRepository interface {
	FindAll(param string, limit int, offset int) (data []model.Task, err error)
	FindByID(id string) (data *model.Task, err error)
	Create(entity model.Task) (data *model.Task, err error)
	Patch(id string, entity model.Task) (data *model.Task, err error)
	Delete(id string) (err error)
}

type TaskService interface {
	FindAll(param string, limit int, offset int) (data []model.Task, err error)
	FindByID(id string) (data *model.Task, err error)
	Create(request request.CreateTaskRequest) (data *model.Task, err error)
	Patch(id string, request request.CreateTaskRequest) (data *model.Task, err error)
	Delete(id string) (err error)
}
