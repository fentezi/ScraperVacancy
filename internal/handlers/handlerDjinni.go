package handlers

import (
	"context"
	"github.com/fentezi/scraper/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetDjinni(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()
	experience := c.Query("experience")
	response, responseErr := services.ParserDjinni(ctx, h.logger, experience)
	if responseErr != nil {
		h.logger.Info(responseErr)
		c.AbortWithError(http.StatusBadRequest, responseErr)
		return
	}
	select {
	case <-ctx.Done():
		return
	case response := <-response:
		c.HTML(http.StatusOK, "index.html", gin.H{
			"infoDjinni": response,
		})
		return
	}
}
