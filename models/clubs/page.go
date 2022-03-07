package clubs

import (
	"database/sql"
	"time"
)

// ClubPage DB モデル
type ClubPage struct {
	ClubUUID         string           `gorm:"type:char(36);not null;primaryKey"`
	ClubSlug         string           `gorm:"type:char(15);not null;unique"`
	Name             string           `gorm:"type:varchar(63);not null"`
	Description      string           `gorm:"type:text;not null"`
	ShortDescription string           `gorm:"type:varchar(50);not null"`
	Campus           uint8            `gorm:"type:tinyint;not null"`
	ClubType         uint8            `gorm:"type:tinyint;not null"`
	ClubRemark       sql.NullString   `gorm:"type:text"`
	ScheduleRemark   sql.NullString   `gorm:"type:text"`
	Visible          bool             `gorm:"type:tinyint(1);not null"`
	UpdatedAt        time.Time        `gorm:"type:datetime;not null"`
	Thumbnail        ClubThumbnail    `gorm:"foreignKey:ClubUUID;references:ClubUUID"`
	Contents         Contents         `gorm:"foreignKey:ClubUUID;references:ClubUUID"`
	Links            Links            `gorm:"foreignKey:ClubUUID;references:ClubUUID"`
	Schedules        Schedules        `gorm:"foreignKey:ClubUUID;references:ClubUUID"`
	Achievements     Achievements     `gorm:"foreignKey:ClubUUID;references:ClubUUID"`
	Images           Images           `gorm:"foreignKey:ClubUUID;references:ClubUUID"`
	Videos           Videos           `gorm:"foreignKey:ClubUUID;references:ClubUUID"`
	ActivityDetails  []ActivityDetail `gorm:"foreignKey:ClubUUID;references:ClubUUID"`
}

type Pages []ClubPage

func (p Pages) GetUUIDs() []string {
	uuids := make([]string, len(p))

	for i, page := range p {
		uuids[i] = page.ClubUUID
	}

	return uuids
}

func (p Pages) ToExternalInfo() []ClubPageExternalInfo {
	exInfo := make([]ClubPageExternalInfo, len(p))

	for i, page := range p {
		info := ClubPageExternalInfo{
			ClubUUID:         page.ClubUUID,
			ClubSlug:         page.ClubSlug,
			Name:             page.Name,
			Description:      page.Description,
			ShortDescription: page.ShortDescription,
			Campus:           page.Campus,
			ClubType:         page.ClubType,
			UpdatedAt:        page.UpdatedAt,
			Thumbnail:        page.Thumbnail.ToThumbnailResponse(),
		}
		exInfo[i] = info
	}

	return exInfo
}

type ClubPageExternalInfo struct {
	ClubUUID         string            `json:"club_uuid"`
	ClubSlug         string            `json:"club_slug"`
	Name             string            `json:"name"`
	Description      string            `json:"description"`
	ShortDescription string            `json:"short_description"`
	Campus           uint8             `json:"campus"`
	ClubType         uint8             `json:"club_type"`
	UpdatedAt        time.Time         `json:"updated_at"`
	Thumbnail        ThumbnailResponse `json:"thumbnail"`
}

type ClubPageInternalInfo struct {
	ClubUUID         string                   `json:"club_uuid"`
	Name             string                   `json:"name"`
	Description      string                   `json:"description"`
	ShortDescription string                   `json:"short_description"`
	Campus           uint8                    `json:"campus"`
	ClubType         uint8                    `json:"club_type"`
	ClubRemark       *string                  `json:"club_remark"`
	ScheduleRemark   *string                  `json:"schedule_remark"`
	UpdatedAt        time.Time                `json:"updated_at"`
	Contents         []ContentResponse        `json:"contents"`
	Links            []LinkResponse           `json:"links"`
	Schedules        []ScheduleResponse       `json:"schedules"`
	Achievements     []AchievementResponse    `json:"achievements"`
	Images           []ImageResponse          `json:"images"`
	Videos           []VideoResponse          `json:"videos"`
	TimePlaces       []ActivityDetailResponse `json:"time_places"`
}
