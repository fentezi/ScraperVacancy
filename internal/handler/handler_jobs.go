package handler

import (
	"context"
	"github.com/fentezi/scraper/pkg/scraper"
	"github.com/gin-gonic/gin"
	"net/http"
	"slices"
	"strings"
	"sync"
)

func (h *handler) ErrorJobs(c *gin.Context) {
	h.logger.Error("Search request missing 'websites' parameter")
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func (h *handler) Jobs(c *gin.Context) {
	var websitesSlice []string
	experience := c.Query("experience")
	websites := c.Query("websites")
	if len(websites) == 0 {
		h.ErrorJobs(c)
		return
	} else {
		websitesSlice = strings.Split(websites, ",")
	}
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()
	h.logger.Info("Job search started")
	chDjinni := make(chan []scraper.InfoDjinni, 1)
	chDou := make(chan []scraper.InfoDou, 1)
	var wg sync.WaitGroup
	switch {
	case slices.Contains(websitesSlice, "dou.ua") && slices.Contains(websitesSlice, "djinni.com"):
		wg.Add(2)
		go scraper.ParserDou(ctx, h.logger, chDou, experience, &wg)
		go scraper.ParseDjinni(ctx, h.logger, chDjinni, experience, &wg)
		h.logger.Debug("Scraping both Djinni and Dou concurrently")

	case slices.Contains(websitesSlice, "dou.ua"):
		wg.Add(1)
		go scraper.ParserDou(ctx, h.logger, chDou, experience, &wg)
		close(chDjinni)
	case slices.Contains(websitesSlice, "djinni.com"):
		wg.Add(1)
		go scraper.ParseDjinni(ctx, h.logger, chDjinni, experience, &wg)
		close(chDou)
	default:
		h.ErrorJobs(c)
		return
	}
	wg.Wait()
	infoDjinni, okDjinni := <-chDjinni
	infoDou, okDou := <-chDou
	select {
	case <-ctx.Done():
		h.logger.Warn("Job search timed out")
		c.AbortWithError(http.StatusRequestTimeout, ctx.Err())
		return
	default:
		switch {
		case okDjinni && okDou:
			c.HTML(http.StatusOK, "index.html", gin.H{
				"infoDjinni": infoDjinni,
				"infoDou":    infoDou,
			})
		case okDjinni:
			c.HTML(http.StatusOK, "index.html", gin.H{
				"infoDjinni": infoDjinni,
			})
		case okDou:
			c.HTML(http.StatusOK, "index.html", gin.H{
				"infoDou": infoDou,
			})
		}
	}
	h.logger.Infof("Job search completed for experience: %s, websites: %v", experience, websites)
}
