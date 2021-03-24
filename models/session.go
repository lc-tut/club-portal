package models

import "github.com/google/uuid"

type AuthState struct {
	UUID  string
	Email string
}

func NewAuthState(uuid uuid.UUID, email string) *AuthState {
	u := uuid.String()
	s := &AuthState{
		UUID:  u,
		Email: email,
	}
	return s
}
