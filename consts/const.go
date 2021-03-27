package consts

const (
	SessionCookieName   = "cp_sess"
	AuthCSRFCookieName  = "cp_auth_csrf"
	SessionKey          = "auth_state"
	CookiePath          = "/"
	CookieHttpOnly      = true
	CookieSessionMaxAge = 60 * 60 * 24 * 7
	CookieCSRFMaxAge    = 60 * 15
)
