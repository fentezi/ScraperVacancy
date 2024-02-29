package handlers

import (
	"github.com/fentezi/scraper/pkg/logging"
	"github.com/gin-gonic/gin"
)

type handler struct {
	logger logging.Logger
}

func NewHandler(logger logging.Logger) *handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) InitRouters() *gin.Engine {
	r := gin.New()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", h.Home)
	r.GET("/jobs/djinni.com", h.GetDjinni)
	return r
}
