package api

import (
	"github.com/go-chi/render"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"net/http"
)

type Error struct {
	Err            error `json:"-"`    // low-level runtime error
	HTTPStatusCode int   `json:"code"` // http response status code

	StatusText string `json:"status"`             // user-level status message
	AppCode    int64  `json:"app_code,omitempty"` // application-specific error code
	ErrorText  string `json:"error,omitempty"`    // application-level error message, for debugging
}

func (e *Error) Render(w http.ResponseWriter, r *http.Request) error {
	log.Ctx(r.Context()).Error().Msgf("error: %+v", errors.Wrap(e.Err, e.ErrorText))
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrBadRequest(err error) render.Renderer {
	return &Error{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "bad request",
		ErrorText:      err.Error(),
	}
}

func ErrUnauthorized(err error) render.Renderer {
	return &Error{
		Err:            err,
		HTTPStatusCode: 403,
		StatusText:     "unauthorized",
		ErrorText:      err.Error(),
	}
}

func ErrNotFound(err error) render.Renderer {
	return &Error{
		Err:            err,
		HTTPStatusCode: 404,
		StatusText:     "resource not found",
		ErrorText:      err.Error(),
	}
}
