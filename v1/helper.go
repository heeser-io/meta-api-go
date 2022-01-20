package v1

import (
	"bytes"
	"encoding/json"
	"io"
)

func structToReader(v interface{}) io.Reader {
	b, err := json.Marshal(v)
	if err != nil {
		return nil
	}
	return bytes.NewReader(b)
}
