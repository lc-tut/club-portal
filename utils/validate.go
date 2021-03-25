package utils

import (
	"github.com/goccy/go-json"
	"github.com/lc-tut/club-portal/models"
)

func ByteSliceToAuthState(b []byte) (*models.AuthState, error) {
	model := &models.AuthState{}
	if err := json.Unmarshal(b, model); err != nil {
		return nil, err
	}

	return model, nil
}
