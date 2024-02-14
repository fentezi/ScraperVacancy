package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRouters() *gin.Engine {
	r := gin.New()
	r.GET("/jobs", h.Jobs)
	return r
}
