package common

const (
	DefaultLimit    int    = 5
	ImagePath       string = "assets/complaints"
	StatusPending   uint   = 0
	StatusOnReceive uint   = 1
	StatusOnProcess uint   = 2
	StatusFinish    uint   = 3
)

type APIResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}