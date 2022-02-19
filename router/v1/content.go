package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/consts"
	"github.com/lc-tut/club-portal/models/clubs"
	"net/http"
)

func (h *Handler) GetClubContent() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clubUUID := ctx.GetString(consts.ClubUUIDKeyName)
		contents, err := h.repo.GetContentsByClubUUID(clubUUID)

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.JSON(http.StatusOK, clubs.Contents(contents).ToContentResponse())
		}
	}
}

func (h *Handler) UpdateClubContent() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var pd []clubs.ContentRequest

		if err := ctx.ShouldBindJSON(pd); err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}

		clubUUID := ctx.GetString(consts.ClubUUIDKeyName)
		conts, err := h.repo.UpdateContent(clubUUID, validateToContentArgs(pd))

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.JSON(http.StatusCreated, clubs.Contents(conts).ToContentResponse())
		}
	}
}
