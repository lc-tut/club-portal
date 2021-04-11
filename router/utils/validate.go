package utils

import (
	"errors"
	"github.com/goccy/go-json"
	"github.com/lc-tut/club-portal/consts"
)

func ByteSliceToSessionData(b []byte) (*SessionData, error) {
	model := &SessionData{}

	if err := json.Unmarshal(b, model); err != nil {
		return nil, err
	}

	return model, nil
}

func ToCampusType(i uint8) (consts.CampusType, error) {
	typed := consts.CampusType(i)
	if typed > consts.CampusHachioji {
		return 0, errors.New("invalid argument")
	} else {
		return typed, nil
	}
}

func ToClubType(i uint8) (consts.ClubType, error) {
	typed := consts.ClubType(i)
	if typed > consts.KokasaiType {
		return 0, errors.New("invalid argument")
	} else {
		return typed, nil
	}
}

func ToVisibility(i uint8) (consts.Visibility, error) {
	typed := consts.Visibility(i)
	if typed > consts.Invisible {
		return 0, errors.New("invalid argument")
	} else {
		return typed, nil
	}
}

func NilToEmptyString(s *string) string {
	if s == nil {
		return ""
	} else {
		return *s
	}
}
