package consts

const (
	SessionCookieName   = "cp_sess"
	AuthCSRFCookieName  = "cp_auth_csrf"
	SessionKey          = "auth_state"
	CookiePath          = "/"
	CookieHttpOnly      = true
	CookieSessionMaxAge = 60 * 60 * 24 * 7
	CookieCSRFMaxAge    = 60 * 15
	UserEmail           = "email_address"
	ClubIDKeyName       = "club_id"
)

// DB のキャンパスタイプ
type CampusType uint

const (
	CampusKamata   CampusType = 0
	CampusHachioji CampusType = 1
)

// DB のサークルタイプ
type ClubType uint

const (
	SportsType  ClubType = 0
	CultureType ClubType = 1
	KokasaiType ClubType = 2
)

// DB の公開か非公開か
type Visibility uint

const (
	Visible   Visibility = 0
	Invisible Visibility = 1
)
