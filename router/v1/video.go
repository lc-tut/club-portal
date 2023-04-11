package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/consts"
	"github.com/lc-tut/club-portal/models/clubs"
	"github.com/lc-tut/club-portal/router/utils"
	"net/http"
)

func (h *Handler) GetClubVideo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clubUUID := ctx.GetString(consts.ClubUUIDKeyName)
		videos, err := h.repo.GetVideosByClubUUID(clubUUID)

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.JSON(http.StatusOK, clubs.Videos(videos).ToVideoResponse())
		}
	}
}

func (h *Handler) UpdateClubVideo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var pd []clubs.VideoRequest

		if err := ctx.ShouldBindJSON(&pd); err != nil {
			h.logger.Error(err.Error())
			ctx.Status(http.StatusBadRequest)
			return
		}

		clubUUID := ctx.GetString(consts.ClubUUIDKeyName)

		if err := h.repo.UpdateVideo(clubUUID, utils.ValidateToVideoArgs(pd)); err != nil {
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.JSON(http.StatusCreated, pd)
		}
	}
}
