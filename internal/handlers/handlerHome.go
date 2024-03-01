package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) Home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", "")
}
