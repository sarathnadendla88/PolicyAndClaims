package errors

import "net/http"

type RestErr struct {
	Message string `json:"message,omitempty"`
	Status  int    `json:"status,omitempty"`
	Error   string `json:"error,omitempty"`
}

func NewBadRequest(Message string) *RestErr {
	return &RestErr{
		Message: Message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}
func NewNotFoundRequest(Message string) *RestErr {
	return &RestErr{
		Message: Message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

func NewInternalServerError(Message string) *RestErr {
	return &RestErr{
		Message: Message,
		Status:  http.StatusInternalServerError,
		Error:   "Internal server error",
	}
}
