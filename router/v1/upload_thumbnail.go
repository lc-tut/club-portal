package v1

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/consts"
	"github.com/lc-tut/club-portal/utils"
	"go.uber.org/zap"
	"image"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func (h *Handler) GetClubThumbnail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clubUUID := ctx.GetString(consts.ClubUUIDKeyName)

		thumbnail, err := h.repo.GetClubThumbnailByUUID(clubUUID)

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.JSON(http.StatusOK, thumbnail.ToThumbnailResponse())
		}
	}
}

func (h *Handler) UploadClubThumbnail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		file, err := ctx.FormFile("file")

		if err != nil {
			h.logger.Error(err.Error())
			ctx.Status(http.StatusBadRequest)
			return
		}

		if err := h.checkThumbnail(file); err != nil {
			h.logger.Error(err.Error())
			ctx.Status(http.StatusBadRequest)
			return
		}

		fn := filepath.Base(file.Filename)
		h.logger.Info("uploaded image", zap.String("filename", fn))
		newFn, err := utils.GenerateFileName(fn)

		if err != nil {
			h.logger.Error(err.Error())
			ctx.Status(http.StatusInternalServerError)
			return
		}

		dst := fmt.Sprintf("thumbnails/%s", newFn)

		if err := ctx.SaveUploadedFile(file, dst); err != nil {
			h.logger.Error(err.Error())
			ctx.Status(http.StatusInternalServerError)
			return
		}

		thumbnail, err := h.repo.CreateThumbnail(dst)

		if err != nil {
			_ = h.deleteSavedThumbnail(dst)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		clubUUID := ctx.GetString(consts.SessionUserUUID) // FIXME: cannot change thumbnail if user is admin

		if err := h.repo.CreateClubThumbnail(clubUUID, thumbnail.ThumbnailID); err != nil {
			_ = h.deleteSavedThumbnail(dst)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		h.logger.Info("successfully saved image", zap.String("path", dst))

		ctx.Status(http.StatusCreated)
	}
}

func (h *Handler) checkThumbnail(file *multipart.FileHeader) error {
	thumbnail, err := file.Open()

	if err != nil {
		return err
	}

	im, _, err := image.Decode(thumbnail)

	if err != nil {
		return err
	}

	rect := im.Bounds()

	if rect.Dx() != 400 && rect.Dy() != 400 {
		return errors.New("thumbnail must be 400x400px size")
	}

	if err := thumbnail.Close(); err != nil {
		return err
	}

	return nil
}

func (h *Handler) deleteSavedThumbnail(file string) error {
	if err := os.Remove(file); err != nil {
		h.logger.Error(err.Error())
		return err
	}

	return nil
}

func (h *Handler) UpdateClubThumbnail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clubUUID := ctx.GetString(consts.ClubUUIDKeyName)

		oldThumbnail, err := h.repo.GetClubThumbnailByUUID(clubUUID)

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return
		}

		file, err := ctx.FormFile("file")

		if err != nil {
			h.logger.Error(err.Error())
			ctx.Status(http.StatusBadRequest)
			return
		}

		if err := h.checkThumbnail(file); err != nil {
			h.logger.Error(err.Error())
			ctx.Status(http.StatusBadRequest)
			return
		}

		fn := filepath.Base(file.Filename)
		h.logger.Info("uploaded image", zap.String("filename", fn))
		newFn, err := utils.GenerateFileName(fn)

		if err != nil {
			h.logger.Error(err.Error())
			ctx.Status(http.StatusInternalServerError)
			return
		}

		dst := fmt.Sprintf("thumbnails/%s", newFn)

		if err := ctx.SaveUploadedFile(file, dst); err != nil {
			h.logger.Error(err.Error())
			ctx.Status(http.StatusInternalServerError)
			return
		}

		if err := h.repo.UpdateThumbnail(oldThumbnail.ThumbnailID, dst); err != nil {
			_ = h.deleteSavedThumbnail(dst)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		if err := h.deleteSavedThumbnail(oldThumbnail.Path); err != nil {
			ctx.Status(http.StatusInternalServerError)
			return
		}

		ctx.Status(http.StatusCreated)
	}
}

func (h *Handler) DeleteClubThumbnail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clubUUID := ctx.GetString(consts.ClubUUIDKeyName)

		oldThumbnail, err := h.repo.GetClubThumbnailByUUID(clubUUID)

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return
		}

		if err := h.repo.UpdateClubThumbnail(clubUUID, 1); err != nil {
			ctx.Status(http.StatusInternalServerError)
			return
		}

		if err := h.repo.DeleteThumbnail(oldThumbnail.ThumbnailID); err != nil {
			ctx.Status(http.StatusInternalServerError)
			return
		}

		if err := h.deleteSavedThumbnail(oldThumbnail.Path); err != nil {
			ctx.Status(http.StatusInternalServerError)
			return
		}

		ctx.Status(http.StatusCreated)
	}
}
