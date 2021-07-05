package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lc-tut/club-portal/consts"
	models "github.com/lc-tut/club-portal/models/users"
	repos "github.com/lc-tut/club-portal/repos/users"
	"net/http"
)

func (h *Handler) GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userUUID := ctx.GetString(consts.SessionUserUUID)
		email := ctx.GetString(consts.SessionUserEmail)
		name := ctx.GetString(consts.SessionUserName)
		role := ctx.GetString(consts.SessionUserRole)

		res := &models.UserResponse{
			UserUUID: userUUID,
			Email:    email,
			Name:     name,
			Role:     role,
		}

		ctx.JSON(http.StatusOK, res)
	}
}

func (h *Handler) GetUserUUID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userUUID := ctx.GetString(consts.UserUUIDKeyName)
		role := ctx.GetString(consts.SessionUserRole) // role が Admin の場合↓が正常に動かない可能性
		user, err := h.repo.GetUserByUUIDFromRole(userUUID, role)

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.JSON(http.StatusOK, user.ToUserResponse())
		}
	}
}

type GUserCreatePostData struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func (h *Handler) CreateGeneralUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pd := &GUserCreatePostData{}

		if err := ctx.ShouldBindJSON(pd); err != nil {
			h.logger.Error(err.Error())
			ctx.Status(http.StatusBadRequest)
			return
		}

		userUUID, err := uuid.NewRandom()

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return
		}

		res, err := h.repo.CreateGeneralUser(userUUID.String(), pd.Email, pd.Name)

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
		} else {
			h.config.WhitelistUsers.AddGeneralUser(res.GetEmail())
			ctx.JSON(http.StatusCreated, res.ToUserResponse())
		}
	}
}

type UserUpdatePostData struct {
	Name     string  `json:"name"`
	Role     string  `json:"role"`
	ClubUUID *string `json:"club_uuid"`
}

func (h *Handler) UpdateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pd := &UserUpdatePostData{}

		if err := ctx.ShouldBindJSON(pd); err != nil {
			h.logger.Error(err.Error())
			ctx.Status(http.StatusBadRequest)
			return
		}

		userUUID := ctx.GetString(consts.UserUUIDKeyName)

		args := repos.UpdateUserArgs{
			Name:     pd.Name,
			ClubUUID: pd.ClubUUID,
		}

		tx := h.repo.UpdateUserFromRole(userUUID, pd.Role, args)

		if err := tx.Error; err != nil {
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.Status(http.StatusCreated)
		}
	}
}
