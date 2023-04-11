package admins

import (
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/consts"
	models "github.com/lc-tut/club-portal/models/clubs"
	adminrepos "github.com/lc-tut/club-portal/repos/admins"
	repos "github.com/lc-tut/club-portal/repos/clubs"
	routerutils "github.com/lc-tut/club-portal/router/utils"
	"github.com/lc-tut/club-portal/utils"
	"net/http"
)

type UpdatePostData struct {
	Visible          bool                           `json:"visible"`
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

func (h *handler) UpdateClubWithVisible() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pd := &UpdatePostData{}

		pageArgs, err := h.makeUpdateArgs(ctx, pd)

		if err != nil {
			h.logger.Error(err.Error())
			ctx.Status(http.StatusBadRequest)
			return
		}

		clubUUID := ctx.GetString(consts.ClubUUIDKeyName)

		if err := h.repo.UpdatePageByClubUUIDWithAdmin(clubUUID, *pageArgs); err != nil {
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.JSON(http.StatusCreated, pd)
		}
	}
}

func (*handler) makeUpdateArgs(ctx *gin.Context, pd *UpdatePostData) (*adminrepos.ClubPageUpdateArgsWithAdmin, error) {
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

	args := &adminrepos.ClubPageUpdateArgsWithAdmin{
		ClubPageUpdateArgs: pageArgs,
		Visible:            pd.Visible,
	}

	return args, nil
}
