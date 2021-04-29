package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/utils"
	"go.uber.org/zap"
	"net/http"
	"path/filepath"
)

func (h *Handler) UploadImage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		form, err := ctx.MultipartForm()

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return
		}

		files := form.File["images"]

		var isError bool

		for _, image := range files {
			filename := filepath.Base(image.Filename)
			h.logger.Info("uploaded image", zap.String("filename", filename))
			newFn, err := utils.GenerateFileName(filename)

			if err != nil {
				h.logger.Error(err.Error())
				isError = true
				break
			}

			if err := ctx.SaveUploadedFile(image, fmt.Sprintf("images/%s", newFn)); err != nil {
				h.logger.Error(err.Error())
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
