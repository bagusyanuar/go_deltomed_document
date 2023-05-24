package request

type CreateProductionSubStepRequest struct {
	ProductionStepID string `json:"production_step_id"`
	Name             string `json:"name"`
	IndexOf          uint   `json:"index_of"`
}
