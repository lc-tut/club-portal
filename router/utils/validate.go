package utils

import (
	"github.com/goccy/go-json"
	"github.com/lc-tut/club-portal/router/data"
)

func ByteSliceToSessionData(b []byte) (*data.SessionData, error) {
	model := &data.SessionData{}

	if err := json.Unmarshal(b, model); err != nil {
		return nil, err
	}

	return model, nil
}
