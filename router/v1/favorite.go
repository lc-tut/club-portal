package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/consts"
	"net/http"
)

func (h *Handler) GetFavoriteClubs() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userUUID := ctx.GetString(consts.UserUUIDKeyName)
		favs, err := h.repo.GetFavorites(userUUID)

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.JSON(http.StatusOK, favs)
		}
	}
}

type IsFavResponse struct {
	Status bool `json:"status"`
}

func (h *Handler) GetIsFavoriteClub() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userUUID := ctx.GetString(consts.UserUUIDKeyName)
		clubUUID := ctx.GetString(consts.ClubUUIDKeyName)
		favs, err := h.repo.GetFavorites(userUUID)

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
		} else {
			for _, club := range favs {
				if club.ClubUUID == clubUUID {
					ctx.JSON(http.StatusOK, &IsFavResponse{Status: true})
					return
				}
			}
			ctx.JSON(http.StatusOK, &IsFavResponse{})
		}
	}
}

type FavPostData struct {
	ClubUUID string `json:"club_uuid"`
}

func (h *Handler) CreateFavoriteClub() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pd := &FavPostData{}
		if err := ctx.ShouldBindJSON(pd); err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}

		userUUID := ctx.GetString(consts.UserUUIDKeyName)

		if err := h.repo.CreateFavorite(userUUID, pd.ClubUUID); err != nil {
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.Status(http.StatusCreated)
		}
	}
}

type UnFavPostData struct {
	ClubUUID string `json:"club_uuid"`
}

func (h *Handler) UnFavoriteClub() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pd := &UnFavPostData{}
		if err := ctx.ShouldBindJSON(pd); err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}

		userUUID := ctx.GetString(consts.UserUUIDKeyName)

		if err := h.repo.DeleteFavorite(userUUID, pd.ClubUUID); err != nil {
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.Status(http.StatusCreated)
		}
	}
}
