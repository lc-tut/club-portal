package utils

type SessionData struct {
	SessionUUID string `json:"session_uuid"`
	UserUUID    string `json:"user_uuid"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	Role        string `json:"role"`
}

func NewSessionData(sessionUUID string, userUUID string, email string, name string, role string) *SessionData {
	s := &SessionData{
		SessionUUID: sessionUUID,
		UserUUID:    userUUID,
		Email:       email,
		Name:        name,
		Role:        role,
	}
	return s
}
