package admins

import (
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/consts"
	"net/http"
)

type UpdateDomainUserPostData struct {
	Name string `json:"name"`
}

func (h *handler) GetUserFromAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userUUID := ctx.GetString(consts.UserUUIDKeyName)
		user, err := h.repo.GetSpecifiedUser(userUUID)

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.JSON(http.StatusOK, user.ToUserResponse())
		}
	}
}

func (h *handler) UpdateDomainUserFromAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pd := &UpdateDomainUserPostData{}

		if err := ctx.ShouldBindJSON(pd); err != nil {
			h.logger.Error(err.Error())
			ctx.Status(http.StatusBadRequest)
			return
		}

		userUUID := ctx.GetString(consts.UserUUIDKeyName)

		if err := h.repo.UpdateSpecifiedDomainUser(userUUID, pd.Name); err != nil {
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.Status(http.StatusCreated)
		}
	}
}
