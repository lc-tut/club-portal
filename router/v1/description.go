package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/consts"
	"net/http"
)

func (h *Handler) GetClubDescription() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clubUUID := ctx.GetString(consts.ClubUUIDKeyName)
		desc, err := h.repo.GetClubDescription(clubUUID)

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.JSON(http.StatusOK, desc)
		}
	}
}

type ClubDescriptionPostData struct {
	Description string `json:"description"`
}

func (h *Handler) UpdateClubDescription() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pd := &ClubCreatePostData{}

		if err := ctx.ShouldBindJSON(pd); err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}

		clubUUID := ctx.GetString(consts.ClubUUIDKeyName)
		desc, err := h.repo.UpdateClubDescription(clubUUID, pd.Description)

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.JSON(http.StatusCreated, desc)
		}
	}
}
