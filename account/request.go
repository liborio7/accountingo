package account

import (
	"errors"
	"net/http"
)

type Request struct {
	*Model
}

func (r *Request) Bind(*http.Request) error {
	// post-processing after a request is unmarshalled
	// from request to model/bean
	if r.Model == nil {
		return errors.New("missing required User fields")
	}
	return nil
}
