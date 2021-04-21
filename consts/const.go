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
	UserUUID            = "user_uuid"
)

// CampusType サークルのキャンパスタイプ (0: 蒲田, 1: 八王子)
type CampusType uint8

func (ct CampusType) ToPrimitive() uint8 {
	return uint8(ct)
}

const (
	CampusKamata   CampusType = 0
	CampusHachioji CampusType = 1
)

// ClubType サークルの種類 (0: 体育会系, 1: 文化会系, 2: 実行委員会)
type ClubType uint8

func (ct ClubType) ToPrimitive() uint8 {
	return uint8(ct)
}

const (
	SportsType  ClubType = 0
	CultureType ClubType = 1
	KokasaiType ClubType = 2
)
