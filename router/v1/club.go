package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lc-tut/club-portal/consts"
	models "github.com/lc-tut/club-portal/models/clubs"
	repos "github.com/lc-tut/club-portal/repos/clubs"
	routerutils "github.com/lc-tut/club-portal/router/utils"
	"github.com/lc-tut/club-portal/utils"
	"net/http"
)

func (h *Handler) GetAllClub() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pages, err := h.repo.GetAllPages()

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.JSON(http.StatusOK, pages)
		}
	}
}

func (h *Handler) GetClubFromSlug() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clubSlug := ctx.GetString(consts.ClubSlugKeyName)
		if ctx.GetBool(consts.IsRestricted) {
			page, err := h.repo.GetRestrictedPageByClubSlug(clubSlug)
			if err != nil {
				ctx.Status(http.StatusInternalServerError)
			} else {
				ctx.JSON(http.StatusOK, page)
			}
		} else {
			page, err := h.repo.GetPageByClubSlug(clubSlug)
			if err != nil {
				ctx.Status(http.StatusInternalServerError)
			} else {
				ctx.JSON(http.StatusOK, page)
			}
		}
	}
}

func (h *Handler) GetClubFromUUID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clubUUID := ctx.GetString(consts.ClubUUIDKeyName)
		if ctx.GetBool(consts.IsRestricted) {
			page, err := h.repo.GetRestrictedPageByClubUUID(clubUUID)
			if err != nil {
				ctx.Status(http.StatusInternalServerError)
			} else {
				ctx.JSON(http.StatusOK, page)
			}
		} else {
			page, err := h.repo.GetPageByClubUUID(clubUUID)
			if err != nil {
				ctx.Status(http.StatusInternalServerError)
			} else {
				ctx.JSON(http.StatusOK, page)
			}
		}
	}
}

type ClubCreatePostData struct {
	Name             string                         `json:"name"`
	Description      string                         `json:"description"`
	ShortDescription string                         `json:"short_description"`
	Campus           uint8                          `json:"campus"`
	ClubType         uint8                          `json:"club_type"`
	ClubRemark       *string                        `json:"club_remark"`
	ScheduleRemark   *string                        `json:"schedule_remark"`
	Contents         []models.ContentRequest        `json:"contents"`
	Links            []models.LinkRequest           `json:"links"`
	Schedules        []models.ScheduleRequest       `json:"schedules"`
	Achievements     []models.AchievementRequest    `json:"achievements"`
	Images           []models.ImageRequest          `json:"images"`
	Videos           []models.VideoRequest          `json:"videos"`
	ActivityDetails  []models.ActivityDetailRequest `json:"activity_details"`
}

func (h *Handler) CreateClub() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := h.checkDuplication(ctx); err != nil {
			h.logger.Error(err.Error())
			ctx.Status(http.StatusBadRequest)
			return
		}

		pd := &ClubCreatePostData{}

		pageArgs, err := h.makeCreateArgs(ctx, pd)

		if err != nil {
			h.logger.Error(err.Error())
			ctx.Status(http.StatusBadRequest)
			return
		}

		if err := h.createPage(ctx, *pageArgs); err != nil {
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.JSON(http.StatusCreated, pd)
		}
	}
}

// Check duplication if general users create new page
func (h *Handler) checkDuplication(ctx *gin.Context) error {
	email := ctx.GetString(consts.SessionUserEmail)

	if h.config.WhitelistUsers.IsGeneralUser(email) {
		userUUID := ctx.GetString(consts.SessionUserUUID)
		gUserModel, err := h.repo.GetGeneralUserByUUID(userUUID)

		if err != nil {
			return err
		}

		clubUUID := utils.NullStringToStringP(gUserModel.ClubUUID)

		if clubUUID != nil {
			return errors.New("already have a club")
		}
	}

	return nil
}

func (*Handler) makeCreateArgs(ctx *gin.Context, pd *ClubCreatePostData) (*repos.ClubPageCreateArgs, error) {
	if err := ctx.ShouldBindJSON(pd); err != nil {
		return nil, err
	}

	campus, err := utils.ToCampusType(pd.Campus)

	if err != nil {
		return nil, err
	}

	clubType, err := utils.ToClubType(pd.ClubType)

	if err != nil {
		return nil, err
	}

	pageArgs := &repos.ClubPageCreateArgs{
		Name:            pd.Name,
		Desc:            pd.Description,
		ShortDesc:       pd.ShortDescription,
		ClubRemark:      utils.StringPToString(pd.ClubRemark),
		ScheduleRemark:  utils.StringPToString(pd.ScheduleRemark),
		Campus:          campus,
		ClubType:        clubType,
		Visible:         true,
		Contents:        routerutils.ValidateToContentArgs(pd.Contents),
		Links:           routerutils.ValidateToLinksArgs(pd.Links),
		Schedules:       routerutils.ValidateToScheduleArgs(pd.Schedules),
		Achievements:    routerutils.ValidateToAchievementArgs(pd.Achievements),
		Images:          routerutils.ValidateToImageArgs(pd.Images),
		Videos:          routerutils.ValidateToVideoArgs(pd.Videos),
		Times:           routerutils.ValidateToTimeArgs(pd.ActivityDetails),
		Places:          routerutils.ValidateToPlaceArgs(pd.ActivityDetails),
		TPRemark:        routerutils.ValidateToTPRemarkArgs(pd.ActivityDetails),
		ActivityDetails: routerutils.ValidateToActivityDetailArgs(pd.ActivityDetails),
	}

	return pageArgs, nil
}

func (h *Handler) createPage(ctx *gin.Context, args repos.ClubPageCreateArgs) error {
	clubUUID, err := uuid.NewRandom()

	if err != nil {
		return err
	}

	page, err := h.repo.CreatePage(clubUUID.String(), args)

	if err != nil {
		return err
	}

	email := ctx.GetString(consts.SessionUserEmail)

	if h.config.WhitelistUsers.IsGeneralUser(email) {
		userUUID := ctx.GetString(consts.SessionUserUUID)
		if err := h.repo.UpdateGeneralUser(userUUID, "", page.ClubUUID); err != nil {
			return err
		}
	}

	return nil
}

type UpdatePostData struct {
	Description      string                         `json:"description"`
	ShortDescription string                         `json:"short_description"`
	ClubRemark       *string                        `json:"club_remark"`
	ScheduleRemark   *string                        `json:"schedule_remark"`
	Contents         []models.ContentRequest        `json:"contents"`
	Links            []models.LinkRequest           `json:"links"`
	Schedules        []models.ScheduleRequest       `json:"schedules"`
	Achievements     []models.AchievementRequest    `json:"achievements"`
	Images           []models.ImageRequest          `json:"images"`
	Videos           []models.VideoRequest          `json:"videos"`
	ActivityDetails  []models.ActivityDetailRequest `json:"activity_details"`
}

func (h *Handler) UpdateClub() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pd := &UpdatePostData{}

		pageArgs, err := h.makeUpdateArgs(ctx, pd)

		if err != nil {
			h.logger.Error(err.Error())
			ctx.Status(http.StatusBadRequest)
			return
		}

		clubUUID := ctx.GetString(consts.ClubUUIDKeyName)

		if err := h.repo.UpdatePageByClubUUID(clubUUID, *pageArgs); err != nil {
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.JSON(http.StatusCreated, pd)
		}
	}
}

func (*Handler) makeUpdateArgs(ctx *gin.Context, pd *UpdatePostData) (*repos.ClubPageUpdateArgs, error) {
	if err := ctx.ShouldBindJSON(pd); err != nil {
		return nil, err
	}

	pageArgs := &repos.ClubPageUpdateArgs{
		Desc:            pd.Description,
		ShortDesc:       pd.ShortDescription,
		ClubRemark:      utils.StringPToString(pd.ClubRemark),
		ScheduleRemark:  utils.StringPToString(pd.ScheduleRemark),
		Contents:        routerutils.ValidateToContentArgs(pd.Contents),
		Links:           routerutils.ValidateToLinksArgs(pd.Links),
		Schedules:       routerutils.ValidateToScheduleArgs(pd.Schedules),
		Achievements:    routerutils.ValidateToAchievementArgs(pd.Achievements),
		Images:          routerutils.ValidateToImageArgs(pd.Images),
		Videos:          routerutils.ValidateToVideoArgs(pd.Videos),
		Times:           routerutils.ValidateToTimeArgs(pd.ActivityDetails),
		Places:          routerutils.ValidateToPlaceArgs(pd.ActivityDetails),
		TPRemark:        routerutils.ValidateToTPRemarkArgs(pd.ActivityDetails),
		ActivityDetails: routerutils.ValidateToActivityDetailArgs(pd.ActivityDetails),
	}

	return pageArgs, nil
}

func (h *Handler) DeleteClub() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clubUUID := ctx.GetString(consts.ClubUUIDKeyName)

		if err := h.repo.DeletePageByClubUUID(clubUUID); err != nil {
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.Status(http.StatusCreated)
		}
	}
}
