package utils

import (
	"database/sql"
	"errors"
	"github.com/lc-tut/club-portal/consts"
)

func StringToNullString(s string) sql.NullString {
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

func StringPToString(s *string) string {
	if s == nil {
		return ""
	} else {
		return *s
	}
}

func NullStringToStringP(s sql.NullString) *string {
	if s.Valid {
		return &s.String
	} else {
		return nil
	}
}

func ToCampusType(i uint8) (consts.CampusType, error) {
	typed := consts.CampusType(i)
	switch typed {
	case consts.CampusKamata, consts.CampusHachioji:
		return typed, nil
	default:
		return 0, nil
	}
}

func ToClubType(i uint8) (consts.ClubType, error) {
	typed := consts.ClubType(i)
	switch typed {
	case consts.SportsType, consts.CultureType, consts.KokasaiType:
		return typed, nil
	default:
		return 0, errors.New("invalid ClubType")
	}
}

func ToUserType(s string) (consts.UserType, error) {
	typed := consts.UserType(s)
	switch typed {
	case consts.AdminUser, consts.GeneralUser, consts.DomainUser:
		return typed, nil
	default:
		return "", errors.New("no role: " + s)
	}
}
