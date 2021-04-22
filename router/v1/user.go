package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/consts"
	"github.com/lc-tut/club-portal/models"
	"net/http"
)

func (h *Handler) GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uuid := ctx.GetString(consts.SessionUserUUID)
		email := ctx.GetString(consts.SessionUserEmail)
		name := ctx.GetString(consts.SessionUserName)
		role := ctx.GetString(consts.SessionUserRole)

		res := &models.UserResponse{
			UserUUID: uuid,
			Email:    email,
			Name:     name,
			Role:     role,
		}

		ctx.JSON(http.StatusOK, res)
	}
}

func (h *Handler) GetUserUUID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uuid := ctx.GetString(consts.UserUUIDKeyName)
		role := ctx.GetString(consts.SessionUserRole)
		user, err := h.repo.GetUserByUUIDFromRole(uuid, role)

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.JSON(http.StatusOK, user.ToUserResponse())
		}
	}
}
