package handler

import (
	"github.com/fentezi/scraper/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) Jobs(c *gin.Context) {
	infoDjinni := pkg.ParseDjinni()
	infoDou := pkg.ParserDou()

	c.HTML(http.StatusOK, "index.html", gin.H{
		"infoDjinni": infoDjinni,
		"infoDou":    infoDou,
	})
}
