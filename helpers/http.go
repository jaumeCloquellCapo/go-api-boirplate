package helpers

import (
	"encoding/json"
	"net/http"
)
type ok interface {
	OK() error
}

type ErrMissingField string

func (e ErrMissingField) Error() string {
	return string(e) + "is required"
}

func DecodeHTTP(r *http.Request, v ok) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	return v.OK()
}
