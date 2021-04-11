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
	ClubSlugKeyName     = "club_slug"
)

// サークルのキャンパスタイプ
type CampusType uint8

func (ct CampusType) ToPrimitive() uint8 {
	return uint8(ct)
}

const (
	CampusKamata   CampusType = 0
	CampusHachioji CampusType = 1
)

// サークルの種類
type ClubType uint8

func (ct ClubType) ToPrimitive() uint8 {
	return uint8(ct)
}

const (
	SportsType  ClubType = 0
	CultureType ClubType = 1
	KokasaiType ClubType = 2
)

// サークルが公開か非公開か
type Visibility uint8

func (v Visibility) ToPrimitive() uint8 {
	return uint8(v)
}

const (
	Visible   Visibility = 0
	Invisible Visibility = 1
)
