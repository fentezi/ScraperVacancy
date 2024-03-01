package handlers

import (
	"context"
	"github.com/fentezi/scraper/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetVacancies(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()
	experience := c.Query("experience")
	responseDjinni, responseErr := services.ParserDjinni(ctx, h.logger, experience)
	if responseErr != nil {
		c.AbortWithError(http.StatusBadRequest, responseErr)
	}
	responseDou, responseErr := services.ParserDjinni(ctx, h.logger, experience)
	if responseErr != nil {
		c.AbortWithError(http.StatusBadRequest, responseErr)
	}
	select {
	case <-ctx.Done():
		c.AbortWithError(http.StatusInternalServerError, ctx.Err())

	case responseDjinni := <-responseDjinni:
		responseDou := <-responseDou
		c.HTML(http.StatusOK, "index.html", gin.H{
			"infoDou":    responseDou,
			"infoDjinni": responseDjinni,
		})
	}
}
