package request

type CreateProductionStepRequest struct {
	ProductionID string `json:"production_id"`
	Name         string `json:"name"`
	IndexOf      uint   `json:"index_of"`
}
