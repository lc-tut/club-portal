package admins

import (
	"github.com/lc-tut/club-portal/models/clubs"
	repos "github.com/lc-tut/club-portal/repos/clubs"
	"github.com/lc-tut/club-portal/utils"
	"gorm.io/gorm"
)

type ClubPageUpdateArgsWithAdmin struct {
	*repos.ClubPageUpdateArgs
	Visible bool
}

type AdminClubRepo interface {
	UpdatePageByClubUUIDWithAdmin(uuid string, args ClubPageUpdateArgsWithAdmin) error
}

func (r *AdminRepository) UpdatePageByClubUUIDWithAdmin(uuid string, args ClubPageUpdateArgsWithAdmin) error {
	page := clubs.ClubPage{
		Visible:          args.Visible,
		Description:      args.Desc,
		ShortDescription: args.ShortDesc,
		ClubRemark:       utils.StringToNullString(args.ClubRemark),
		ScheduleRemark:   utils.StringToNullString(args.ScheduleRemark),
	}

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&page).Where("club_uuid = ?", uuid).Select("Visible", "Description", "ShortDescription", "ClubRemark", "ScheduleRemark").Updates(page).Error; err != nil {
			return err
		}

		if err := r.UpdateContentWithTx(tx, uuid, args.Contents); err != nil {
			return err
		}

		if err := r.UpdateLinkWithTx(tx, uuid, args.Links); err != nil {
			return err
		}

		if err := r.UpdateScheduleWithTx(tx, uuid, args.Schedules); err != nil {
			return err
		}

		if err := r.UpdateAchievementWithTx(tx, uuid, args.Achievements); err != nil {
			return err
		}

		if err := r.UpdateImageWithTx(tx, uuid, args.Images); err != nil {
			return err
		}

		if err := r.UpdateVideoWithTx(tx, uuid, args.Videos); err != nil {
			return err
		}

		if err := r.CreateTimeWithTx(tx, args.Times); err != nil {
			return err
		}

		if err := r.CreatePlaceWithTx(tx, args.Places); err != nil {
			return err
		}

		if err := r.UpdateActivityDetailWithTx(tx, uuid, args.ActivityDetails); err != nil {
			return err
		}

		if err := r.UpdateTPRemarkWithTx(tx, uuid, args.TPRemark); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
