package consts

const (
	SessionCookieName     = "cp_sess"
	AuthCSRFCookieName    = "cp_auth_csrf"
	RedirectURLCookieName = "cp_redirect"
	SessionKey            = "auth_state"
	CookiePath            = "/"
	CookieHttpOnly        = true
	CookieSessionMaxAge   = 60 * 60 * 24 * 7
	CookieCSRFMaxAge      = 60 * 15
	SessionUserEmail      = "sess_user_email"
	SessionUserUUID       = "sess_user_uuid"
	SessionUserName       = "sess_user_name"
	SessionUserRole       = "sess_user_role"
	IsRestricted          = "is_restricted"
	ClubSlugKeyName       = "club_slug_key"
	UserUUIDKeyName       = "user_uuid_key"
	ImageIDKeyName        = "image_id_key"
	ClubUUIDKeyName       = "club_uuid_key"
	ThumbnailIDKeyName    = "thumbnail_id_key"
)

// CampusType サークルのキャンパスタイプ (0: 八王子, 1: 蒲田)
type CampusType uint8

const (
	CampusHachioji CampusType = 0
	CampusKamata   CampusType = 1
)

// ToPrimitive CampusType を uint8 に変換
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

// ToPrimitive ClubType を uint8 に変換
func (ct ClubType) ToPrimitive() uint8 {
	return uint8(ct)
}

// UserType ユーザの種類 (domain: 学内ユーザ, general: 一般ユーザ, admin: 管理者)
type UserType string

const (
	DomainUser  UserType = "domain"
	GeneralUser UserType = "general"
	AdminUser   UserType = "admin"
)

// ToPrimitive UserType を string に変換
func (ut UserType) ToPrimitive() string {
	return string(ut)
}

// UploadSize アップロード可能なファイルの最大サイズ (2MiB)
const UploadSize = 2 << 18 // 2MiB

// DummyUUID ダミーのUUID
const DummyUUID = "aaaaaaaa-aaaa-4aaa-aaaa-aaaaaaaaaaaa"

// DefaultThumbnailPath デフォルトのサムネイル画像のパス
const DefaultThumbnailPath = "thumbnails/default.png"
