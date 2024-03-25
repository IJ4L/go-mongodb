package domain

type SuccessResponse struct {
	Status  bool      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}