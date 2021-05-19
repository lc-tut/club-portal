package clubs

import (
	"github.com/lc-tut/club-portal/consts"
	"github.com/lc-tut/club-portal/models/clubs"
	"github.com/lc-tut/club-portal/utils"
	"gorm.io/gorm"
)

type ClubPageCreateArgs struct {
	Name            string
	Desc            string
	Campus          consts.CampusType
	ClubType        consts.ClubType
	Visible         bool
	Contents        []string
	Links           []ClubLinkArgs
	Schedules       []ClubScheduleArgs
	Achievements    []string
	Images          []uint32
	Videos          []string
	Times           []ClubTimeArgs
	Places          []ClubPlaceArgs
	Remarks         []ClubRemarkArgs
	ActivityDetails []ActivityDetailArgs
}

type ClubPageUpdateArgs struct {
	Desc            string
	Contents        []string
	Links           []ClubLinkArgs
	Schedules       []ClubScheduleArgs
	Achievements    []string
	Images          []uint32
	Videos          []string
	Times           []ClubTimeArgs
	Places          []ClubPlaceArgs
	Remarks         []ClubRemarkArgs
	ActivityDetails []ActivityDetailArgs
}

type ClubPageRepo interface {
	GetAllPages() ([]clubs.ClubPageExternalInfo, error)
	GetPageByClubUUID(uuid string) (*clubs.ClubPageInternalInfo, error)
	GetPageByClubSlug(clubSlug string) (*clubs.ClubPageInternalInfo, error)

	CreatePage(uuid string, args ClubPageCreateArgs) (*clubs.ClubPage, error)

	UpdatePageByClubUUID(uuid string, args ClubPageUpdateArgs) error
	UpdatePageByClubSlug(clubSlug string, args ClubPageUpdateArgs) error

	DeletePageByClubUUID(uuid string) error
	DeletePageByClubSlug(slug string) error
}

func (r *ClubRepository) GetAllPages() ([]clubs.ClubPageExternalInfo, error) {
	page := make([]clubs.ClubPage, 0)
	tx := r.db.Where("visible is true").Preload("Thumbnail", func(db *gorm.DB) *gorm.DB {
		selectQuery := "club_thumbnails.thumbnail_id, club_thumbnails.club_uuid, ut.path"
		joinQuery := "inner join uploaded_thumbnails as ut using (thumbnail_id)"
		return db.Joins(joinQuery).Select(selectQuery)
	}).Find(&page)

	if err := tx.Error; err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	typedPage := clubs.Pages(page)

	return typedPage.ToExternalInfo(), nil
}

func (r *ClubRepository) GetPageByClubUUID(uuid string) (*clubs.ClubPageInternalInfo, error) {
	page := &clubs.ClubPage{}
	tx := r.db.Where("club_uuid = ? and visible is true", uuid).Preload("Contents").Preload("Links").Preload("Schedules").Preload("Achievements").Preload("Videos").Preload("ActivityDetails").Take(page)

	if err := tx.Error; err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	info, err := r.getPage(page)

	if err != nil {
		return nil, err
	}

	return info, nil
}

func (r *ClubRepository) GetPageByClubSlug(clubSlug string) (*clubs.ClubPageInternalInfo, error) {
	page := &clubs.ClubPage{}
	tx := r.db.Where("club_slug = ? and visible is true", clubSlug).Preload("Contents").Preload("Links").Preload("Schedules").Preload("Achievements").Preload("Images").Preload("Videos").Preload("ActivityDetails").Take(page)

	if err := tx.Error; err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}

	info, err := r.getPage(page)

	if err != nil {
		return nil, err
	}

	return info, nil
}

func (r *ClubRepository) getPage(page *clubs.ClubPage) (*clubs.ClubPageInternalInfo, error) {
	rels, err := r.GetAllRelations(page.ClubUUID)

	if err != nil {
		return nil, err
	}

	images, err := r.GetImagesByClubUUID(page.ClubUUID)

	if err != nil {
		return nil, err
	}

	typedRels := clubs.Relations(rels)

	typedImages := clubs.Images(images)

	info := &clubs.ClubPageInternalInfo{
		ClubUUID:     page.ClubUUID,
		Name:         page.Name,
		Description:  page.Description,
		Campus:       page.Campus,
		ClubType:     page.ClubType,
		UpdatedAt:    page.UpdatedAt,
		Contents:     page.Contents.ToContentResponse(),
		Links:        page.Links.ToLinkResponse(),
		Schedules:    page.Schedules.ToScheduleResponse(),
		Achievements: page.Achievements.ToAchievementResponse(),
		Images:       typedImages.ToImageResponse(),
		Videos:       page.Videos.ToVideoResponse(),
		Times:        clubs.Times(typedRels.ToClubTime()).ToTimeResponse(typedRels.ToClubRemark()),
		Places:       clubs.Places(typedRels.ToClubPlace()).ToPlaceResponse(typedRels.ToClubRemark()),
	}

	return info, nil
}

func (r *ClubRepository) CreatePage(uuid string, args ClubPageCreateArgs) (*clubs.ClubPage, error) {
	slug := utils.GenerateSlug(uuid)

	page := &clubs.ClubPage{
		ClubUUID:    uuid,
		ClubSlug:    slug,
		Name:        args.Name,
		Description: args.Desc,
		Campus:      args.Campus.ToPrimitive(),
		ClubType:    args.ClubType.ToPrimitive(),
		Visible:     args.Visible,
	}

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(page).Error; err != nil {
			return err
		}

		if err := r.CreateContentWithTx(tx, uuid, args.Contents); err != nil {
			return err
		}

		if err := r.CreateLinkWithTx(tx, uuid, args.Links); err != nil {
			return err
		}

		if err := r.CreateScheduleWithTx(tx, uuid, args.Schedules); err != nil {
			return err
		}

		if err := r.CreateAchievementWithTx(tx, uuid, args.Achievements); err != nil {
			return err
		}

		if err := r.CreateImageWithTx(tx, uuid, args.Images); err != nil {
			return err
		}

		if err := r.CreateVideoWithTx(tx, uuid, args.Videos); err != nil {
			return err
		}

		if err := r.CreateTimeWithTx(tx, args.Times); err != nil {
			return err
		}

		if err := r.CreatePlaceWithTx(tx, args.Places); err != nil {
			return err
		}

		if err := r.CreateActivityDetailWithTx(tx, uuid, args.ActivityDetails); err != nil {
			return err
		}

		if err := r.CreateRemarkWithTx(tx, uuid, args.Remarks); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return page, nil
}

func (r *ClubRepository) UpdatePageByClubUUID(uuid string, args ClubPageUpdateArgs) error {
	page := clubs.ClubPage{
		Description: args.Desc,
	}

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&page).Where("club_uuid = ?", uuid).Updates(page).Error; err != nil {
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

		if err := r.UpdateRemarkWithTx(tx, uuid, args.Remarks); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *ClubRepository) UpdatePageByClubSlug(clubSlug string, args ClubPageUpdateArgs) error {
	page := clubs.ClubPage{}

	tx := r.db.Where("club_slug = ?", clubSlug).Select("club_uuid").Take(&page)

	if err := tx.Error; err != nil {
		r.logger.Error(err.Error())
		return err
	}

	if err := r.UpdatePageByClubUUID(page.ClubUUID, args); err != nil {
		return err
	}

	return nil
}

func (r *ClubRepository) DeletePageByClubUUID(uuid string) error {
	tx := r.db.Model(&clubs.ClubPage{}).Where("club_uuid = ?", uuid).Update("visible", false)

	if err := tx.Error; err != nil {
		r.logger.Error(err.Error())
		return err
	}

	return nil
}

func (r *ClubRepository) DeletePageByClubSlug(slug string) error {
	tx := r.db.Model(&clubs.ClubPage{}).Where("club_slug = ?", slug).Update("visible", false)

	if err := tx.Error; err != nil {
		r.logger.Error(err.Error())
		return err
	}

	return nil
}
