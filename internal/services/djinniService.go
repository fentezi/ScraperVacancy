package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/fentezi/scraper/models"
	"github.com/fentezi/scraper/pkg/logging"
	"github.com/gocolly/colly/v2"
	"strings"
)

func ParserDjinni(ctx context.Context, logger logging.Logger, experience string) (chan []models.Djinni, error) {
	if experience == "" {
		return nil, errors.New("experience is empty")
	}
	if experience != "Middle" && experience != "Senior" && experience != "Junior" {
		return nil, errors.New("choose the right experience")
	}
	col := colly.NewCollector()
	var infoDjinni []models.Djinni

	const baseUrl = "https://djinni.co/jobs/?primary_keyword=Golang"

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		col.OnHTML("div.job-list-item", func(e *colly.HTMLElement) {
			position := strings.TrimSpace(e.ChildText("a.h3.job-list-item__link"))
			if strings.HasPrefix(position, experience) {
				views := e.ChildTexts("span.mr-2")
				info := models.Djinni{
					Href:     "https://djinni.co" + e.ChildAttr("a.h3.job-list-item__link", "href"),
					Date:     e.ChildAttr("span.mr-2.nobr", "title"),
					Position: position,
					Views:    views[2],
					Reviews:  views[3],
				}
				infoDjinni = append(infoDjinni, info)
			}
		})
		currentPage := 1
		col.OnScraped(func(response *colly.Response) {
			logger.Infof("Finisged scraping: %s", response.Request.URL)
			if currentPage < 4 {
				currentPage++
			}
			nextUrl := fmt.Sprintf("%s&page=%d", baseUrl, currentPage)
			err := col.Visit(nextUrl)
			if err != nil {
				return
			}
		})

		col.OnError(func(r *colly.Response, err error) {
			logger.Errorf("Request URL: %s failed with response: %v\nError: %v", r.Request.URL, r, err)
		})

		err := col.Visit(fmt.Sprintf("%s&page=1", baseUrl))
		if err != nil {
			return nil, err
		}
		chDjinni := make(chan []models.Djinni, 1)
		chDjinni <- infoDjinni
		return chDjinni, nil
	}

}
