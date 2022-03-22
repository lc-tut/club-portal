package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type SearchContent struct {
	Content string `form:"content" binding:"required"`
}

func (h *Handler) GetSearch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		searchContent := &SearchContent{}

		if err := ctx.ShouldBind(searchContent); err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}

		pages, err := h.repo.DoSearch(searchContent.Content)

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.JSON(http.StatusOK, pages)
		}
	}
}
