package response

import (
	"net/http"
)

type Writer struct {
	http.ResponseWriter
	status int
}

func (w *Writer) Status() int {
	return w.status
}

func (w *Writer) WriteHeader(statusCode int) {
	w.status = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
