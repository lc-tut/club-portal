package utils

import (
	"github.com/goccy/go-json"
	models "github.com/lc-tut/club-portal/models/users"
)

func ByteSliceToSessionData(b []byte) (*SessionData, error) {
	model := &SessionData{}

	if err := json.Unmarshal(b, model); err != nil {
		return nil, err
	}

	return model, nil
}

func ToUserInfoResponse(users []models.UserInfo) []models.UserResponse {
	res := make([]models.UserResponse, len(users))

	for i, user := range users {
		res[i] = *user.ToUserResponse()
	}

	return res
}
