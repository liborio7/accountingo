package api

import (
	"net/http"
)

type Response struct {
	http.ResponseWriter
	status int
}

func NewResponse(r http.ResponseWriter) *Response {
	return &Response{ResponseWriter: r}
}

func (w *Response) Status() int {
	return w.status
}

func (w *Response) WriteHeader(statusCode int) {
	w.status = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
