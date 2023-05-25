package admin

import (
	"fmt"
	"time"

	"github.com/bagusyanuar/go_deltomed_document/common"
	"github.com/bagusyanuar/go_deltomed_document/http/request"
	"github.com/bagusyanuar/go_deltomed_document/model"
	usecase "github.com/bagusyanuar/go_deltomed_document/usecase/admin"
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type implementsTaskService struct {
	TaskRepository usecase.TaskRepository
}

// Create implements admin.TaskService
func (service *implementsTaskService) Create(request request.CreateTaskRequest) (data *model.Task, err error) {
	productionID, err := uuid.Parse(request.ProductionID)
	if err != nil {
		return nil, err
	}
	supportID, err := uuid.Parse(request.SupportID)
	if err != nil {
		return nil, err
	}

	startDate, err := time.Parse(common.DateFormat, request.StartDate)
	if err != nil {
		return nil, err
	}
	finishDate, err := time.Parse(common.DateFormat, request.FinishDate)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	formattedTime := now.Format(common.CodeDateFormat)
	code := fmt.Sprintf("TASK-%s", formattedTime)
	entity := model.Task{
		ProductionID: productionID,
		SupportID:    supportID,
		Code:         code,
		StartDate:    datatypes.Date(startDate),
		FinishDate:   datatypes.Date(finishDate),
		CompletedAt:  nil,
	}
	return service.TaskRepository.Create(entity)
}

// Delete implements admin.TaskService
func (service *implementsTaskService) Delete(id string) (err error) {
	return service.TaskRepository.Delete(id)
}

// FindAll implements admin.TaskService
func (service *implementsTaskService) FindAll(param string, limit int, offset int) (data []model.Task, err error) {
	if limit == 0 {
		limit = common.DefaultLimit
	}
	return service.TaskRepository.FindAll(param, limit, offset)
}

// FindByID implements admin.TaskService
func (service *implementsTaskService) FindByID(id string) (data *model.Task, err error) {
	return service.TaskRepository.FindByID(id)
}

// Patch implements admin.TaskService
func (service *implementsTaskService) Patch(id string, request request.CreateTaskRequest) (data *model.Task, err error) {
	productionID, err := uuid.Parse(request.ProductionID)
	if err != nil {
		return nil, err
	}
	supportID, err := uuid.Parse(request.SupportID)
	if err != nil {
		return nil, err
	}

	startDate, err := time.Parse(common.DateFormat, request.StartDate)
	if err != nil {
		return nil, err
	}
	finishDate, err := time.Parse(common.DateFormat, request.FinishDate)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	formattedTime := now.Format(common.CodeDateFormat)
	code := fmt.Sprintf("TASK-%s", formattedTime)
	entity := model.Task{
		ProductionID: productionID,
		SupportID:    supportID,
		Code:         code,
		StartDate:    datatypes.Date(startDate),
		FinishDate:   datatypes.Date(finishDate),
		CompletedAt:  nil,
	}
	return service.TaskRepository.Patch(id, entity)
}

func NewTaskService(taskRepository usecase.TaskRepository) usecase.TaskService {
	return &implementsTaskService{
		TaskRepository: taskRepository,
	}
}
