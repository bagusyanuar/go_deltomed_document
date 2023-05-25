package request

type CreateTaskProcessRequest struct {
	TaskID              string `json:"task_id"`
	ProductionStepID    string `json:"production_step_id"`
	ProductionSubStepID string `json:"production_sub_step_id"`
	Image               string `json:"image"`
}
