package account

import (
	"net/http"
)

type Response struct {
	*Model
}

func NewResponse(model *Model) *Response {
	// from model/bean to response
	return &Response{Model: model}
}

func (r *Response) Render(http.ResponseWriter, *http.Request) error {
	// pre-processing before a response is marshalled
	return nil
}

type PaginatedResponse struct {
	Data    []*Response
	HasMore bool
}

func NewPaginatedResponse(models []Model, limit int) *PaginatedResponse {
	var data []*Response
	var length int
	if len(models) < limit {
		length = len(models)
	} else {
		length = limit
	}
	for i := range models[:length] {
		data = append(data, NewResponse(&models[i]))
	}
	return &PaginatedResponse{Data: data, HasMore: len(models) > limit}
}

func (r *PaginatedResponse) Render(http.ResponseWriter, *http.Request) error {
	// pre-processing before a response is marshalled
	return nil
}
