package support

import "github.com/bagusyanuar/go_deltomed_document/model"

type TaskProcessRepository interface {
	GetDataByTaskID(authorizedID string, id string) (data []model.TaskProcess, err error)
}

type TaskProcessService interface {
	GetDataByTaskID(authorizedID string, id string) (data []model.TaskProcess, err error)
}
