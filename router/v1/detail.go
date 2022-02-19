package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/consts"
	"github.com/lc-tut/club-portal/models/clubs"
	"net/http"
)

type response struct {
	Times  []clubs.TimeResponse  `json:"times"`
	Places []clubs.PlaceResponse `json:"places"`
}

func (h *Handler) GetClubActivityDetails() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clubUUID := ctx.GetString(consts.ClubUUIDKeyName)
		details, err := h.repo.GetAllRelations(clubUUID)

		rels := clubs.Relations(details)

		res := &response{
			Times:  clubs.Times(rels.ToClubTime()).ToTimeResponse(rels.ToClubRemark()),
			Places: clubs.Places(rels.ToClubPlace()).ToPlaceResponse(rels.ToClubRemark()),
		}

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.JSON(http.StatusOK, res)
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

		clubUUID := ctx.GetString(consts.ClubUUIDKeyName)
		timeArgs := validateToTimeArgs(pd)
		placeArgs := validateToPlaceArgs(pd)
		detailsArgs := validateToActivityDetailArgs(pd)
		tpremarkArgs := validateToTPRemarkArgs(pd)

		err := h.repo.UpdateAllRelations(clubUUID, timeArgs, placeArgs, detailsArgs, tpremarkArgs)

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.JSON(http.StatusOK, pd)
		}
	}
}
