package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/consts"
	"github.com/lc-tut/club-portal/models/clubs"
	"github.com/lc-tut/club-portal/router/utils"
	"net/http"
)

func (h *Handler) GetClubActivityDetails() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clubUUID := ctx.GetString(consts.ClubUUIDKeyName)
		details, err := h.repo.GetAllRelations(clubUUID)

		rels := clubs.Relations(details)

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.JSON(http.StatusOK, rels.ToActivityDetailResponse())
		}
	}
}

func (h *Handler) UpdateClubActivityDetails() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var pd []clubs.ActivityDetailRequest

		if err := ctx.ShouldBindJSON(&pd); err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}

		if len(pd) == 0 {
			ctx.Status(http.StatusBadRequest)
			return
		}

		clubUUID := ctx.GetString(consts.ClubUUIDKeyName)
		timeArgs := utils.ValidateToTimeArgs(pd)
		placeArgs := utils.ValidateToPlaceArgs(pd)
		detailsArgs := utils.ValidateToActivityDetailArgs(pd)
		tpremarkArgs := utils.ValidateToTPRemarkArgs(pd)

		err := h.repo.UpdateAllRelations(clubUUID, timeArgs, placeArgs, detailsArgs, tpremarkArgs)

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.JSON(http.StatusCreated, pd)
		}
	}
}
