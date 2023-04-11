package admins

import (
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/consts"
	"github.com/lc-tut/club-portal/repos/admins"
	"github.com/lc-tut/club-portal/router/utils"
	"net/http"
)

func (h *handler) GetAllUserFromAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		users, err := h.repo.GetAllUser()

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.JSON(http.StatusOK, utils.ToUserInfoResponse(users))
		}
	}
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

type UpdateUserPostData struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	ClubUUID string `json:"club_uuid,omitempty"`
}

func (h *handler) UpdateUserFromAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pd := &UpdateUserPostData{}

		if err := ctx.ShouldBindJSON(pd); err != nil {
			h.logger.Error(err.Error())
			ctx.Status(http.StatusBadRequest)
			return
		}

		userUUID := ctx.GetString(consts.UserUUIDKeyName)
		user, err := h.repo.GetSpecifiedUser(userUUID)

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
		} else {
			switch user.GetRole() {
			case consts.DomainUser:
				if err := h.repo.UpdateSpecifiedDomainUser(userUUID, pd.Name); err != nil {
					ctx.Status(http.StatusInternalServerError)
				} else {
					ctx.Status(http.StatusCreated)
				}
			case consts.GeneralUser:
				args := admins.UserArgs{
					Email:    pd.Email,
					Name:     pd.Name,
					ClubUUID: pd.ClubUUID,
				}
				if err := h.repo.UpdateSpecifiedGeneralUser(userUUID, args); err != nil {
					ctx.Status(http.StatusInternalServerError)
				} else {
					ctx.Status(http.StatusCreated)
				}
			}
		}
	}
}
