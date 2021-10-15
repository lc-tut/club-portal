package utils

type SessionData struct {
	SessionUUID string `json:"session_uuid"`
	UserUUID    string `json:"user_uuid"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	Role        string `json:"role"`
	AvatarURL   string `json:"avatar"`
}

func NewSessionData(sessionUUID, userUUID, email, name, role, avatarURL string) *SessionData {
	s := &SessionData{
		SessionUUID: sessionUUID,
		UserUUID:    userUUID,
		Email:       email,
		Name:        name,
		Role:        role,
		AvatarURL:   avatarURL,
	}
	return s
}
