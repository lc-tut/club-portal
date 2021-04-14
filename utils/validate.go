package utils

import (
	"database/sql"
	"errors"
	"github.com/lc-tut/club-portal/consts"
)

func ToNullString(s string) sql.NullString {
	var ns sql.NullString

	if s == "" {
		ns = sql.NullString{
			String: "",
			Valid:  false,
		}
	} else {
		ns = sql.NullString{
			String: s,
			Valid:  true,
		}
	}

	return ns
}

func NilToEmptyString(s *string) string {
	if s == nil {
		return ""
	} else {
		return *s
	}
}

func ToNilOrString(s sql.NullString) *string {
	if s.Valid {
		return &s.String
	} else {
		return nil
	}
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
