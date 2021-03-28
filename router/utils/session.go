package utils

import "github.com/google/uuid"

type SessionData struct {
	UUID  string
	Email string
}

func NewSessionData(uuid uuid.UUID, email string) *SessionData {
	u := uuid.String()
	s := &SessionData{
		UUID:  u,
		Email: email,
	}
	return s
}
