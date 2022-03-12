package v1

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/consts"
	"github.com/lc-tut/club-portal/models/users"
	"github.com/lc-tut/club-portal/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"image"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func (h *Handler) GetImages() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userUUID := ctx.GetString(consts.SessionUserUUID)

		im, err := h.repo.GetImagesByUserUUID(userUUID)

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.JSON(http.StatusOK, users.Images(im).ToImageResponse())
		}
	}
}

func (h *Handler) UploadImage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		form, err := ctx.MultipartForm()

		if err != nil {
			h.logger.Error(err.Error())
			ctx.Status(http.StatusInternalServerError)
			return
		}

		files, ok := form.File["images"]

		if !ok {
			h.logger.Error("name of form should be `images`")
			ctx.Status(http.StatusBadRequest)
			return
		}

		if err := h.checkImage(files); err != nil {
			h.logger.Error(err.Error())
			ctx.Status(http.StatusBadRequest)
			return
		}

		var isError bool

		userUUID := ctx.GetString(consts.SessionUserUUID)

		res := make([]users.ImageResponse, len(files))

		for i, f := range files {
			filename := filepath.Base(f.Filename)
			h.logger.Info("uploaded image", zap.String("filename", filename), zap.String("user_uuid", userUUID))
			newFn, err := utils.GenerateFileName(filename)

			if err != nil {
				h.logger.Error(err.Error())
				isError = true
				break
			}

			dst := fmt.Sprintf("images/%s", newFn)

			if err := ctx.SaveUploadedFile(f, dst); err != nil {
				h.logger.Error(err.Error())
				isError = true
				break
			}

			repoRes, err := h.repo.CreateUploadedImage(userUUID, dst)

			if err != nil {
				_ = h.deleteSavedImage(dst)
				isError = true
				break
			}

			res[i] = users.ImageResponse{
				ImageID: repoRes.ImageID,
				Path:    dst,
			}

			h.logger.Info("successfully saved image", zap.String("path", dst), zap.String("user_uuid", userUUID))
		}

		if isError {
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.JSON(http.StatusCreated, res)
		}
	}
}

func (h *Handler) checkImage(files []*multipart.FileHeader) error {
	for _, f := range files {
		im, err := f.Open()

		if err != nil {
			return err
		}

		_, _, err = image.Decode(im)

		if err != nil {
			return err
		}

		if err := im.Close(); err != nil {
			return err
		}
	}

	return nil
}

func (h *Handler) deleteSavedImage(file string) error {
	if err := os.Remove(file); err != nil {
		h.logger.Error(err.Error())
		return err
	}

	return nil
}

func (h *Handler) GetSpecificImage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		imageID := ctx.GetUint64(consts.ImageIDKeyName)

		im, err := h.repo.GetUploadedImageByID(uint32(imageID))

		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.Status(http.StatusBadRequest)
		} else if err != nil {
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.JSON(http.StatusOK, im)
		}
	}
}

func (h *Handler) DeleteImage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		imageID := ctx.GetUint(consts.ImageIDKeyName)

		im, err := h.repo.GetUploadedImageByID(uint32(imageID))

		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.Status(http.StatusBadRequest)
			return
		} else if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return
		}

		if err := h.repo.DeleteImageByID(uint32(imageID)); err != nil {
			ctx.Status(http.StatusInternalServerError)
			return
		}

		if err := h.deleteSavedImage(im.Path); err != nil {
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.Status(http.StatusOK)
		}
	}
}
