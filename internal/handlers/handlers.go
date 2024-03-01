package handlers

import (
	"github.com/fentezi/scraper/pkg/logging"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	logger logging.Logger
}

func NewHandler(logger logging.Logger) *Handler {
	return &Handler{
		logger: logger,
	}
}

func (h *Handler) InitRouters() *gin.Engine {
	r := gin.New()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", h.Home)
	r.GET("/jobs/djinni.com", h.GetDjinni)
	r.GET("/jobs/dou.ua", h.GetDou)
	r.GET("/jobs/dou.ua,djinni.com", h.GetVacancies)
	return r
}
