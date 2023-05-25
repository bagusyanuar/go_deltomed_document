package request

type CreateTaskRequest struct {
	ProductionID string `json:"production_id"`
	SupportID    string `json:"support_id"`
	Name         string `json:"name"`
	StartDate    string `json:"start_date"`
	FinishDate   string `json:"finish_date"`
}
