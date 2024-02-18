package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *handler) Home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", "")
}
