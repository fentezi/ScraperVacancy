package handlers

import (
	"context"
	"github.com/fentezi/scraper/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetDou(c *gin.Context) {
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()
	experience := c.Query("experience")
	response, responseErr := services.ParserDou(ctx, h.logger, experience)
	if responseErr != nil {
		c.AbortWithError(http.StatusBadRequest, responseErr)
		return
	}
	select {
	case <-ctx.Done():
		c.AbortWithError(http.StatusInternalServerError, ctx.Err())

	case response := <-response:
		c.HTML(http.StatusOK, "index.html", gin.H{
			"infoDou": response,
		})
	}
}
