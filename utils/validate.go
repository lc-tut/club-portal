package utils

import "database/sql"

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

func ValidateIDValue(id uint32, inc *uint32) uint32 {
	if id == 0 {
		id = *inc
		*inc++
	}

	return id
}
