package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lc-tut/club-portal/consts"
	"github.com/lc-tut/club-portal/models"
	"github.com/lc-tut/club-portal/repos"
	"github.com/lc-tut/club-portal/router/utils"
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
	Name         string                       `json:"name"`
	Description  string                       `json:"description"`
	Campus       uint8                        `json:"campus"`
	ClubType     uint8                        `json:"club_type"`
	Visible      uint8                        `json:"visible"`
	Contents     []models.ContentRequest      `json:"contents"`
	Links        []models.LinkRequest         `json:"links"`
	Schedules    []models.ScheduleRequest     `json:"schedules"`
	Achievements *[]models.AchievementRequest `json:"achievements"`
	Images       *[]models.ImageRequest       `json:"images"`
	Videos       *[]models.VideoRequest       `json:"videos"`
	Times        []models.TimeRequest         `json:"times"`
	Places       []models.PlaceRequest        `json:"places"`
}

func (h *Handler) CreateClub() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pd := &CreatePostData{}

		pageArgs, err := h.createArgs(ctx, pd)

		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}

		if err := h.createPage(pd, *pageArgs); err != nil {
			ctx.Status(http.StatusInternalServerError)
			return
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
		Name:     pd.Name,
		Desc:     pd.Description,
		Campus:   campus,
		ClubType: clubType,
		Visible:  visible,
	}

	return pageArgs, nil
}

func (h *Handler) createPage(pd *CreatePostData, pa repos.ClubPageCreateArgs) error {
	clubUUID, err := uuid.NewRandom()

	if err != nil {
		return err
	}

	cu := clubUUID.String()

	if err := h.repo.CreatePage(cu, pa); err != nil {
		return err
	}

	if err := h.repo.CreateContent(cu, ValidateToContentArgs(pd.Contents)); err != nil {
		return err
	}

	if err := h.repo.CreateLink(cu, ValidateToLinksArgs(pd.Links)); err != nil {
		return err
	}

	if err := h.repo.CreateSchedule(cu, ValidateToScheduleArgs(pd.Schedules)); err != nil {
		return err
	}

	if err := h.repo.CreateAchievement(cu, ValidateToAchievementArgs(pd.Achievements)); err != nil {
		return err
	}

	if err := h.repo.CreateImage(cu, ValidateToImageArgs(pd.Images)); err != nil {
		return err
	}

	if err := h.repo.CreateVideo(cu, ValidateToVideoArgs(pd.Videos)); err != nil {
		return err
	}

	if err := h.repo.CreateClubTimeAndPlaces(cu, pd.TimesAndPlaces); err != nil {
		return err
	}

	return nil
}
