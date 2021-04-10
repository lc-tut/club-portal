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
