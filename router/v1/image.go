package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/consts"
	"github.com/lc-tut/club-portal/models/clubs"
	"net/http"
)

func (h *Handler) GetClubImages() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clubUUID := ctx.GetString(consts.ClubUUIDKeyName)
		images, err := h.repo.GetImagesByClubUUID(clubUUID)

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.JSON(http.StatusOK, clubs.Images(images).ToImageResponse())
		}
	}
}

func (h *Handler) UpdateClubImages() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var pd []clubs.ImageRequest

		if err := ctx.ShouldBindJSON(&pd); err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}

		clubUUID := ctx.GetString(consts.ClubUUIDKeyName)

		if err := h.repo.UpdateImage(clubUUID, validateToImageArgs(pd)); err != nil {
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.JSON(http.StatusCreated, pd)
		}
	}
}
