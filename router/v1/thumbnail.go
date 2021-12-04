package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/consts"
	"net/http"
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
