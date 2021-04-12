package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lc-tut/club-portal/consts"
	"github.com/lc-tut/club-portal/models"
	"github.com/lc-tut/club-portal/repos"
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

func (h *Handler) GetClub() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clubSlug := ctx.GetString(consts.ClubSlugKeyName)
		page, err := h.repo.GetPageByClubSlug(clubSlug)

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.JSON(http.StatusOK, page)
		}
	}
}

type CreatePostData struct {
	Name            string                         `json:"name"`
	Description     string                         `json:"description"`
	Campus          uint8                          `json:"campus"`
	ClubType        uint8                          `json:"club_type"`
	Visible         uint8                          `json:"visible"`
	Contents        []models.ContentRequest        `json:"contents"`
	Links           []models.LinkRequest           `json:"links"`
	Schedules       []models.ScheduleRequest       `json:"schedules"`
	Achievements    *[]models.AchievementRequest   `json:"achievements"`
	Images          *[]models.ImageRequest         `json:"images"`
	Videos          *[]models.VideoRequest         `json:"videos"`
	ActivityDetails []models.ActivityDetailRequest `json:"activity_details"`
}

func (h *Handler) CreateClub() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pd := &CreatePostData{}

		pageArgs, err := h.createArgs(ctx, pd)

		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}

		if err := h.createPage(*pageArgs); err != nil {
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.JSON(http.StatusCreated, pd)
		}
	}
}

func (*Handler) createArgs(ctx *gin.Context, pd *CreatePostData) (*repos.ClubPageCreateArgs, error) {
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

	visible, err := utils.ToVisibility(pd.Visible)

	if err != nil {
		return nil, err
	}

	pageArgs := &repos.ClubPageCreateArgs{
		Name:            pd.Name,
		Desc:            pd.Description,
		Campus:          campus,
		ClubType:        clubType,
		Visible:         visible,
		Contents:        validateToContentArgs(pd.Contents),
		Links:           validateToLinksArgs(pd.Links),
		Schedules:       validateToScheduleArgs(pd.Schedules),
		Achievements:    validateToAchievementArgs(*pd.Achievements),
		Images:          validateToImageArgs(*pd.Images),
		Videos:          validateToVideoArgs(*pd.Videos),
		Times:           validateToTimeArgs(pd.ActivityDetails),
		Places:          validateToPlaceArgs(pd.ActivityDetails),
		Remarks:         validateToRemarkArgs(pd.ActivityDetails),
		ActivityDetails: validateToActivityDetailArg(pd.ActivityDetails),
	}

	return pageArgs, nil
}

func (h *Handler) createPage(args repos.ClubPageCreateArgs) error {
	clubUUID, err := uuid.NewRandom()

	if err != nil {
		return err
	}

	if err := h.repo.CreatePage(clubUUID.String(), args); err != nil {
		return err
	}

	return nil
}
