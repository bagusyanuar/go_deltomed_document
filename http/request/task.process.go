package request

import "mime/multipart"

type CreateTaskProcessRequest struct {
	TaskID              string                `json:"task_id"`
	ProductionStepID    string                `json:"production_step_id"`
	ProductionSubStepID string                `json:"production_sub_step_id"`
	Image               *multipart.FileHeader `json:"image"`
}
