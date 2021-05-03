package v1

import (
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

func (h *Handler) UploadImage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		form, err := ctx.MultipartForm()

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return
		}

		files := form.File["images"]

		if err := h.checkImage(files); err != nil {
			h.logger.Error(err.Error())
			ctx.Status(http.StatusBadRequest)
			return
		}

		var isError bool

		userUUID := ctx.GetString(consts.SessionUserName)

		for _, f := range files {
			filename := filepath.Base(f.Filename)
			h.logger.Info("uploaded image", zap.String("filename", filename))
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

			if err := h.repo.CreateUploadedImage(userUUID, dst); err != nil {
				_ = h.deleteImage(dst)
				isError = true
				break
			}
		}

		if isError {
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.Status(http.StatusCreated)
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

func (h *Handler) deleteImage(file string) error {
	if err := os.Remove(file); err != nil {
		h.logger.Error(err.Error())
		return err
	}

	return nil
}
