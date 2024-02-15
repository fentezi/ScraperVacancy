package handler

import (
	"github.com/fentezi/scraper/pkg"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func (h *Handler) Jobs(c *gin.Context) {
	start := time.Now()
	chDjinni := make(chan []pkg.InfoDjinni)
	chDou := make(chan []pkg.InfoDou)
	go func() {
		chDjinni <- pkg.ParseDjinni()
	}()
	go func() {
		chDou <- pkg.ParserDou()
	}()
	infoDjinni := <-chDjinni
	infoDou := <-chDou
	log.Println("Time:", time.Since(start))
	c.HTML(http.StatusOK, "index.html", gin.H{
		"infoDjinni": infoDjinni,
		"infoDou":    infoDou,
	})
}
