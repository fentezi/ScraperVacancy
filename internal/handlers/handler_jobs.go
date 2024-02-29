package handlers

//
//import (
//	"context"
//	"github.com/fentezi/scraper/internal/services"
//	"github.com/fentezi/scraper/pkg/logging"
//	"github.com/gin-gonic/gin"
//	"net/http"
//	"strings"
//	"sync"
//)
//
//func (h *handler) ErrorJobs(c *gin.Context) {
//	h.logger.Error("Search request missing 'websites' parameter")
//	c.Redirect(http.StatusTemporaryRedirect, "/")
//}
//
//func (h *handler) Jobs(c *gin.Context) {
//	var websitesSlice []string
//	ctx, cancel := context.WithCancel(c.Request.Context())
//	defer cancel()
//	experience := c.Query("experience")
//	if experience == "" {
//		h.ErrorJobs(c)
//		return
//	}
//	websites := c.Query("websites")
//	websitesSlice = strings.Split(websites, ",")
//	scrapers := map[string]func(context.Context, logging.Logger, , string, *sync.WaitGroup){
//		"djinni.com": services.ParserDjinni,
//		"dou.ua":     services.ParserDou,
//	}
//	h.logger.Info("Job search started")
//	var wg sync.WaitGroup
//	resultChan := make(map[string]chan interface{})
//	for _, website := range websitesSlice {
//		wg.Add(1)
//		ch := make(chan interface{}, 1)
//		scrp, _ := scrapers[website]
//		go scrp(ctx, h.logger, ch, experience, &wg)
//		resultChan[website] = ch
//	}
//
//	wg.Wait()
//	if len(websitesSlice) == 2 {
//		infoDou := <-resultChan["dou.ua"]
//		infoDjinni := <-resultChan["djinni.com"]
//		c.HTML(http.StatusOK, "index.html", gin.H{
//			"infoDou":    infoDou,
//			"infoDjinni": infoDjinni,
//		})
//		return
//	}
//	select {
//	case <-ctx.Done():
//		h.logger.Info("Canceled")
//		return
//	case infoDjinni := <-resultChan["djinni.com"]:
//		c.HTML(http.StatusOK, "index.html", gin.H{
//			"infoDjinni": infoDjinni,
//		})
//	case infoDou := <-resultChan["dou.ua"]:
//		c.HTML(http.StatusOK, "index.html", gin.H{
//			"infoDou": infoDou,
//		})
//	}
//	h.logger.Infof("Job search completed for experience: %s, websites: %v", experience, websites)
//}
