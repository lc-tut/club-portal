package repos

import (
	"github.com/lc-tut/club-portal/consts"
	"github.com/lc-tut/club-portal/models"
	"github.com/lc-tut/club-portal/utils"
)

type ClubPageCreateArgs struct {
	Name            string
	Desc            string
	Campus          consts.CampusType
	ClubType        consts.ClubType
	Visible         consts.Visibility
	Contents        []string
	Links           []ClubLinkArgs
	Schedules       []ClubScheduleArgs
	Achievements    []string
	Images          []string
	Videos          []string
	Times           []ClubTimeArgs
	Places          []ClubPlaceArgs
	Remarks         []ClubRemarkArgs
	ActivityDetails []ClubActivityDetailArgs
}

type ClubPageUpdateArgs struct {
	Desc    string
	Visible consts.Visibility
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
	tx := r.db.Preload("Contents").Preload("Links").Preload("Schedules").Preload("Achievements").Preload("Images").Preload("Videos").Preload("ActivityDetails").Find(&page)

	if err := tx.Error; err != nil {
		return nil, err
	}

	typedPage := models.Pages(page)

	return typedPage.ToExternalInfo(), nil
}

func (r *Repository) GetPageByClubUUID(uuid string) (*models.ClubPageInternalInfo, error) {
	page := models.ClubPage{}
	tx := r.db.Where("club_uuid = ?", uuid).Preload("Contents").Preload("Links").Preload("Schedules").Preload("Achievements").Preload("Images").Preload("Videos").Preload("ActivityDetails").Take(&page)

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
		Achievements: page.Achievements.ToAchievementResponse(),
		Images:       page.Images.ToImageResponse(),
		Videos:       page.Videos.ToVideoResponse(),
		Times:        models.Times(typedRels.ToClubTime()).ToTimeResponse(typedRels.ToClubRemark()),
		Places:       models.Places(typedRels.ToClubPlace()).ToPlaceResponse(typedRels.ToClubRemark()),
	}

	return info, nil
}

func (r *Repository) GetPageByClubSlug(clubSlug string) (*models.ClubPageInternalInfo, error) {
	page := &models.ClubPage{}
	tx := r.db.Where("club_slug = ?", clubSlug).Preload("Contents").Preload("Links").Preload("Schedules").Preload("Achievements").Preload("Images").Preload("Videos").Preload("ActivityDetails").Take(page)

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
		Achievements: page.Achievements.ToAchievementResponse(),
		Images:       page.Images.ToImageResponse(),
		Videos:       page.Videos.ToVideoResponse(),
		Times:        models.Times(typedRels.ToClubTime()).ToTimeResponse(typedRels.ToClubRemark()),
		Places:       models.Places(typedRels.ToClubPlace()).ToPlaceResponse(typedRels.ToClubRemark()),
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
		Visible:     args.Visible.ToPrimitive(),
	}

	if err := r.db.Create(page).Error; err != nil {
		return err
	}

	if err := r.CreateContent(uuid, args.Contents); err != nil {
		return err
	}

	if err := r.CreateLink(uuid, args.Links); err != nil {
		return err
	}

	if err := r.CreateSchedule(uuid, args.Schedules); err != nil {
		return err
	}

	if err := r.CreateAchievement(uuid, args.Achievements); err != nil {
		return err
	}

	if err := r.CreateImage(uuid, args.Images); err != nil {
		return err
	}

	if err := r.CreateVideo(uuid, args.Videos); err != nil {
		return err
	}

	if err := r.CreateTime(args.Times); err != nil {
		return err
	}

	if err := r.CreatePlace(args.Places); err != nil {
		return err
	}

	if err := r.CreateClubActivityDetail(uuid, args.ActivityDetails); err != nil {
		return err
	}

	if err := r.CreateRemark(uuid, args.Remarks); err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdatePageByClubUUID(uuid string, args ClubPageUpdateArgs) error {
	tx := r.db.Model(&models.ClubPage{}).Where("club_uuid = ?", uuid).Updates(args)

	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdatePageByClubSlug(clubSlug string, args ClubPageUpdateArgs) error {
	tx := r.db.Model(&models.ClubPage{}).Where("club_slug = ?", clubSlug).Updates(args)

	if err := tx.Error; err != nil {
		return err
	}

	return nil
}
