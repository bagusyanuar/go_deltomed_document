package support

import "github.com/bagusyanuar/go_deltomed_document/model"

type TaskRepository interface {
	GetData(authorizedID string, param string, limit int, offset int) (data []model.Task, err error)
}

type TaskService interface {
	GetData(authorizedID string, param string, limit int, offset int) (data []model.Task, err error)
}
