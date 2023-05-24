package request

type CreateTaskRequest struct {
	ProductionID string `json:"production_id"`
	SupportID    string `json:"support_id"`
	Name         string `json:"name"`
	Date         string `json:"date"`
}
