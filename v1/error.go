package v1

import (
	"encoding/json"
	"errors"
)

type HttpError struct {
	Error string `json:"error"`
}

func ErrorFromBody(b []byte) error {
	httpError := &HttpError{}
	if err := json.Unmarshal(b, httpError); err != nil {
		httpError.Error = string(b)
	}

	return errors.New(httpError.Error)
}

func New(msg string) *HttpError {
	return &HttpError{
		Error: msg,
	}
}
func (he *HttpError) String() string {
	res, err := json.Marshal(he)
	if err != nil {
		return "{}"
	}
	return string(res)
}
