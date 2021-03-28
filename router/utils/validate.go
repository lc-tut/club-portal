package utils

import (
	"github.com/goccy/go-json"
)

func ByteSliceToSessionData(b []byte) (*SessionData, error) {
	model := &SessionData{}

	if err := json.Unmarshal(b, model); err != nil {
		return nil, err
	}

	return model, nil
}
