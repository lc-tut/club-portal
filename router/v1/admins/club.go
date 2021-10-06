package admins

import (
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/repos/admins"
	"net/http"
)

type UpdatePostData struct {
	UserUUID string
	Email    string
	Name     string
	ClubUUID string
}

func (h *handler) UpdateClubUserFromAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pd := &UpdatePostData{}

		if err := ctx.ShouldBindJSON(pd); err != nil {
			h.logger.Error(err.Error())
			ctx.Status(http.StatusBadRequest)
			return
		}

		args := admins.UserArgs{
			Email:    pd.Email,
			Name:     pd.Name,
			ClubUUID: pd.ClubUUID,
		}

		if err := h.repo.UpdateSpecifiedClub(pd.UserUUID, args); err != nil {
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.Status(http.StatusCreated)
		}
	}
}
