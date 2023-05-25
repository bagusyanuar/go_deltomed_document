package support

import (
	"github.com/bagusyanuar/go_deltomed_document/common"
	"github.com/bagusyanuar/go_deltomed_document/model"
	usecase "github.com/bagusyanuar/go_deltomed_document/usecase/support"
)

type implementsTaskService struct {
	TaskRepository usecase.TaskRepository
}

// GetData implements support.TaskService
func (service *implementsTaskService) GetData(authorizedID string, param string, limit int, offset int) (data []model.Task, err error) {
	if limit == 0 {
		limit = common.DefaultLimit
	}
	return service.TaskRepository.GetData(authorizedID, param, limit, offset)
}

func NewTaskService(taskRepository usecase.TaskRepository) usecase.TaskService {
	return &implementsTaskService{
		TaskRepository: taskRepository,
	}
}
