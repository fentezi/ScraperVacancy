package scraper

import (
	"context"
	"fmt"
	"github.com/fentezi/scraper/pkg/logging"
	"strings"
	"sync"

	"github.com/gocolly/colly/v2"
)

type InfoDjinni struct {
	Position string `json:"position"`
	Href     string `json:"href"`
	Date     string `json:"date_published"`
	Views    string `json:"views"`
	Reviews  string `json:"reviews"`
}

func ParserDjinni(ctx context.Context, logger logging.Logger, chDjinni chan interface{}, experience string, wg *sync.WaitGroup) {
	col := colly.NewCollector()
	var infoDjinni []InfoDjinni
	defer wg.Done()

	const baseUrl = "https://djinni.co/jobs/?primary_keyword=Golang"

	col.OnHTML("div.job-list-item", func(e *colly.HTMLElement) {
		select {
		case <-ctx.Done():
			return
		default:
			position := strings.TrimSpace(e.ChildText("a.h3.job-list-item__link"))
			if strings.HasPrefix(position, experience) {
				views := e.ChildTexts("span.mr-2")
				info := InfoDjinni{
					Href:     "https://djinni.co" + e.ChildAttr("a.h3.job-list-item__link", "href"),
					Date:     e.ChildAttr("span.mr-2.nobr", "title"),
					Position: position,
					Views:    views[2],
					Reviews:  views[3],
				}
				infoDjinni = append(infoDjinni, info)
			}
		}
	})

	currentPage := 1
	col.OnScraped(func(response *colly.Response) {
		logger.Infof("Finisged scraping: %s", response.Request.URL)
		if currentPage < 4 {
			currentPage++
		}
		nextUrl := fmt.Sprintf("%s&page=%d", baseUrl, currentPage)
		col.Visit(nextUrl)
	})

	col.OnError(func(r *colly.Response, err error) {
		logger.Errorf("Request URL: %s failed with response: %v\nError: %v", r.Request.URL, r, err)
	})

	col.Visit(fmt.Sprintf("%s&page=1", baseUrl))
	chDjinni <- infoDjinni
}
