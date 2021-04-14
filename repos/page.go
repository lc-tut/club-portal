package repos

import (
	"github.com/lc-tut/club-portal/consts"
	"github.com/lc-tut/club-portal/models"
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
	Images          []string
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
	Images          []string
	Videos          []string
	Times           []ClubTimeArgs
	Places          []ClubPlaceArgs
	Remarks         []ClubRemarkArgs
	ActivityDetails []ActivityDetailArgs
}

type ClubPageRepo interface {
	GetAllPages() ([]models.ClubPageExternalInfo, error)
	GetPageByClubUUID(uuid string) (*models.ClubPageInternalInfo, error)
	GetPageByClubSlug(clubSlug string) (*models.ClubPageInternalInfo, error)

	CreatePage(uuid string, args ClubPageCreateArgs) error

	UpdatePageByClubUUID(uuid string, args ClubPageUpdateArgs) error
	UpdatePageByClubSlug(clubSlug string, args ClubPageUpdateArgs) error
}

func (r *Repository) GetAllPages() ([]models.ClubPageExternalInfo, error) {
	page := make([]models.ClubPage, 0)
	tx := r.db.Where("visible is true").Preload("Contents").Preload("Links").Preload("Schedules").Preload("Achievements").Preload("Images").Preload("Videos").Preload("ActivityDetails").Find(&page)

	if err := tx.Error; err != nil {
		return nil, err
	}

	typedPage := models.Pages(page)

	return typedPage.ToExternalInfo(), nil
}

func (r *Repository) GetPageByClubUUID(uuid string) (*models.ClubPageInternalInfo, error) {
	page := models.ClubPage{}
	tx := r.db.Where("club_uuid = ? and visible is true", uuid).Preload("Contents").Preload("Links").Preload("Schedules").Preload("Achievements").Preload("Images").Preload("Videos").Preload("ActivityDetails").Take(&page)

	if err := tx.Error; err != nil {
		return nil, err
	}

	rels, err := r.GetAllRelations(page.ClubUUID)

	if err != nil {
		return nil, err
	}

	typedRels := models.Relations(rels)

	info := &models.ClubPageInternalInfo{
		ClubUUID:     uuid,
		Name:         page.Name,
		Description:  page.Description,
		Campus:       page.Campus,
		ClubType:     page.ClubType,
		UpdatedAt:    page.UpdatedAt,
		Contents:     *page.Contents.ToContentResponse(),
		Links:        *page.Links.ToLinkResponse(),
		Schedules:    *page.Schedules.ToScheduleResponse(),
		Achievements: *page.Achievements.ToAchievementResponse(),
		Images:       *page.Images.ToImageResponse(),
		Videos:       *page.Videos.ToVideoResponse(),
		Times:        *models.Times(typedRels.ToClubTime()).ToTimeResponse(typedRels.ToClubRemark()),
		Places:       *models.Places(typedRels.ToClubPlace()).ToPlaceResponse(typedRels.ToClubRemark()),
	}

	return info, nil
}

func (r *Repository) GetPageByClubSlug(clubSlug string) (*models.ClubPageInternalInfo, error) {
	page := &models.ClubPage{}
	tx := r.db.Where("club_slug = ? and visible is true", clubSlug).Preload("Contents").Preload("Links").Preload("Schedules").Preload("Achievements").Preload("Images").Preload("Videos").Preload("ActivityDetails").Take(page)

	if err := tx.Error; err != nil {
		return nil, err
	}

	rels, err := r.GetAllRelations(page.ClubUUID)

	if err != nil {
		return nil, err
	}

	typedRels := models.Relations(rels)

	info := &models.ClubPageInternalInfo{
		ClubUUID:     page.ClubUUID,
		Name:         page.Name,
		Description:  page.Description,
		Campus:       page.Campus,
		ClubType:     page.ClubType,
		UpdatedAt:    page.UpdatedAt,
		Contents:     *page.Contents.ToContentResponse(),
		Links:        *page.Links.ToLinkResponse(),
		Schedules:    *page.Schedules.ToScheduleResponse(),
		Achievements: *page.Achievements.ToAchievementResponse(),
		Images:       *page.Images.ToImageResponse(),
		Videos:       *page.Videos.ToVideoResponse(),
		Times:        *models.Times(typedRels.ToClubTime()).ToTimeResponse(typedRels.ToClubRemark()),
		Places:       *models.Places(typedRels.ToClubPlace()).ToPlaceResponse(typedRels.ToClubRemark()),
	}

	return info, nil
}

func (r *Repository) CreatePage(uuid string, args ClubPageCreateArgs) error {
	slug, err := utils.GenerateRand15()

	if err != nil {
		return err
	}

	page := &models.ClubPage{
		ClubUUID:    uuid,
		ClubSlug:    slug,
		Name:        args.Name,
		Description: args.Desc,
		Campus:      args.Campus.ToPrimitive(),
		ClubType:    args.ClubType.ToPrimitive(),
		Visible:     args.Visible,
	}

	err = r.db.Transaction(func(tx *gorm.DB) error {
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
		return err
	}

	return nil
}

func (r *Repository) UpdatePageByClubUUID(uuid string, args ClubPageUpdateArgs) error {
	page := models.ClubPage{
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

func (r *Repository) UpdatePageByClubSlug(clubSlug string, args ClubPageUpdateArgs) error {
	page := models.ClubPage{}

	tx := r.db.Where("club_slug = ?", clubSlug).Select("club_uuid").Take(&page)

	if err := tx.Error; err != nil {
		return err
	}

	if err := r.UpdatePageByClubUUID(page.ClubUUID, args); err != nil {
		return err
	}

	return nil
}
