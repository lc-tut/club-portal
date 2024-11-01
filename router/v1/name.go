package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/consts"
)

type ClubNameResponse struct {
	Name string `json:"name"`
}

func (h *Handler) UpdateClubName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pd := &ClubCreatePostData{}

		if err := ctx.ShouldBindJSON(pd); err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}

		clubUUID := ctx.GetString(consts.ClubUUIDKeyName)

		if err := h.repo.UpdateClubName(clubUUID, pd.Name); err != nil {
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.JSON(http.StatusCreated, pd)
		}
	}
}