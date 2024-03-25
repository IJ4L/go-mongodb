package domain

type ErrorResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func NewErrorResponse(err error) ErrorResponse {
	return ErrorResponse{
		Status:  false,
		Message: err.Error(),
	}
}