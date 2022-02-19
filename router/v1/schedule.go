package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/consts"
	"github.com/lc-tut/club-portal/models/clubs"
	"net/http"
)

func (h *Handler) GetClubSchedule() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clubUUID := ctx.GetString(consts.ClubUUIDKeyName)
		schedules, err := h.repo.GetSchedulesByClubUUID(clubUUID)

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.JSON(http.StatusOK, clubs.Schedules(schedules).ToScheduleResponse())
		}
	}
}

func (h *Handler) UpdateClubSchedule() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var pd []clubs.ScheduleRequest

		if err := ctx.ShouldBindJSON(&pd); err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}

		clubUUID := ctx.GetString(consts.ClubUUIDKeyName)

		if err := h.repo.UpdateSchedule(clubUUID, validateToScheduleArgs(pd)); err != nil {
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.JSON(http.StatusOK, pd)
		}
	}
}
