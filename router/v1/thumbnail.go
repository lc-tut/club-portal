package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/consts"
)

func (h *Handler) GetThumbnail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.GetUint(consts.ThumbnailIDKeyName)

		thumbnail, err := h.repo.GetThumbnail(uint32(id))

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.JSON(http.StatusOK, thumbnail.ToThumbnailResponse())
		}
	}
}
