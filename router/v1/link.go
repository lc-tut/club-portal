package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/consts"
	"github.com/lc-tut/club-portal/models/clubs"
	"github.com/lc-tut/club-portal/router/utils"
	"net/http"
)

func (h *Handler) GetClubLinks() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clubUUID := ctx.GetString(consts.ClubUUIDKeyName)
		links, err := h.repo.GetLinksByClubUUID(clubUUID)

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.JSON(http.StatusOK, clubs.Links(links).ToLinkResponse())
		}
	}
}

func (h *Handler) UpdateClubLinks() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var pd []clubs.LinkRequest

		if err := ctx.ShouldBindJSON(&pd); err != nil {
			h.logger.Error(err.Error())
			ctx.Status(http.StatusBadRequest)
			return
		}

		clubUUID := ctx.GetString(consts.ClubUUIDKeyName)

		if err := h.repo.UpdateLink(clubUUID, utils.ValidateToLinksArgs(pd)); err != nil {
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.JSON(http.StatusCreated, pd)
		}
	}
}
