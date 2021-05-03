package consts

const (
	SessionCookieName   = "cp_sess"
	AuthCSRFCookieName  = "cp_auth_csrf"
	SessionKey          = "auth_state"
	CookiePath          = "/"
	CookieHttpOnly      = true
	CookieSessionMaxAge = 60 * 60 * 24 * 7
	CookieCSRFMaxAge    = 60 * 15
	SessionUserEmail    = "sess_user_email"
	SessionUserUUID     = "sess_user_uuid"
	SessionUserName     = "sess_user_name"
	SessionUserRole     = "sess_user_role"
	ClubSlugKeyName     = "club_slug_key"
	UserUUIDKeyName     = "user_uuid_key"
	ImageIDKeyName      = "image_id_key"
)

// CampusType サークルのキャンパスタイプ (0: 蒲田, 1: 八王子)
type CampusType uint8

const (
	CampusKamata   CampusType = 0
	CampusHachioji CampusType = 1
)

func (ct CampusType) ToPrimitive() uint8 {
	return uint8(ct)
}

// ClubType サークルの種類 (0: 体育会系, 1: 文化会系, 2: 実行委員会)
type ClubType uint8

const (
	SportsType  ClubType = 0
	CultureType ClubType = 1
	KokasaiType ClubType = 2
)

func (ct ClubType) ToPrimitive() uint8 {
	return uint8(ct)
}

type UserType string

const (
	DomainUser  UserType = "domain"
	GeneralUser UserType = "general"
	AdminUser   UserType = "admin"
)

func (ut UserType) ToPrimitive() string {
	return string(ut)
}

const UploadSize = 2 << 18 // 2MiB
