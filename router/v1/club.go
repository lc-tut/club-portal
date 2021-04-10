package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/consts"
	"github.com/lc-tut/club-portal/models"
	"net/http"
	"time"
)

type ClubPageExternalInfo struct {
	ClubUUID    string
	ClubSlug    string
	Name        string
	Description string
	Campus      uint8
	ClubType    uint8
	UpdatedAt   time.Time
	Images      []models.ClubImage
}

type ClubPageInternalInfo struct {
	ClubUUID       string
	Name           string
	Description    string
	Campus         uint8
	ClubType       uint8
	UpdatedAt      time.Time
	Contents       []models.ClubContent
	Links          []models.ClubLink
	Schedules      []models.ClubSchedule
	Achievements   []models.ClubAchievement
	Images         []models.ClubImage
	Videos         []models.ClubVideo
	TimesAndPlaces []models.ClubTimeAndPlace
}

func (h *Handler) GetAllClub() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clubs, err := h.repo.GetAllPages()

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return
		}

		res := make([]ClubPageExternalInfo, 0)

		for _, c := range clubs {
			if c.Visible == uint8(consts.Invisible) {
				continue
			} else {
				page := &ClubPageExternalInfo{
					ClubUUID:    c.ClubUUID,
					ClubSlug:    c.ClubSlug,
					Name:        c.Name,
					Description: c.Description,
					Campus:      c.Campus,
					ClubType:    c.ClubType,
					UpdatedAt:   c.UpdatedAt,
					Images:      c.Images,
				}

				res = append(res, *page)
			}
		}

		ctx.JSON(http.StatusOK, res)

	}
}

func (h *Handler) GetClub() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clubSlug := ctx.GetString(consts.ClubSlugKeyName)

		club, err := h.repo.GetPageByClubSlug(clubSlug)

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return
		}

		tps, err := h.repo.GetClubTimeAndPlaces(club.ClubUUID)

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return
		}

		res := &ClubPageInternalInfo{
			ClubUUID:       club.ClubUUID,
			Name:           club.Name,
			Description:    club.Description,
			Campus:         club.Campus,
			ClubType:       club.ClubType,
			UpdatedAt:      club.UpdatedAt,
			Contents:       club.Contents,
			Links:          club.Links,
			Schedules:      club.Schedules,
			Achievements:   club.Achievements,
			Images:         club.Images,
			Videos:         club.Videos,
			TimesAndPlaces: tps,
		}

		ctx.JSON(http.StatusOK, res)
	}
}
