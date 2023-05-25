package common

const (
	DefaultLimit    int    = 5
	ImagePath       string = "assets/complaints"
	StatusPending   uint   = 0
	StatusOnReceive uint   = 1
	StatusOnProcess uint   = 2
	StatusFinish    uint   = 3
	DateFormat      string = "2006-01-02"
	CodeDateFormat  string = "20060102150405"
)

type APIResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
