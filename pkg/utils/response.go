package utils

import (
	"time"
)

func NewResponse(resp Response) Response {
	return Response{
		Status:    resp.Status,
		Message:   resp.Message,
		Data:      resp.Data,
		Timestamp: time.Now(),
	}
}

type Response struct {
	Status    int         `json:"status"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"` // Data is optional, used in success responses
	Timestamp time.Time   `json:"timestamp"`
}
