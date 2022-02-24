package v1

import (
	"encoding/json"
	"errors"
)

type HttpError struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func ErrorFromBody(b []byte) error {
	httpError := &HttpError{}
	if err := json.Unmarshal(b, httpError); err != nil {
		httpError.Error = string(b)
	}
	if httpError.Error == "" {
		return errors.New(httpError.Message)
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
