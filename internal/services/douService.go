package services

import (
	"context"
	"errors"
	"github.com/fentezi/scraper/models"
	"github.com/fentezi/scraper/pkg/logging"
	"github.com/gocolly/colly/v2"
)

func ParserDou(ctx context.Context, logger logging.Logger, experience string) (chan []models.Dou, error) {
	col := colly.NewCollector()

	switch experience {
	case "Junior":
		experience = "0-1"
	case "Middle":
		experience = "1-3"
	case "Senior":
		experience = "3-5"
	default:
		return nil, errors.New("experience is empty")
	}

	col.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36"

	URLDOU := "https://jobs.dou.ua/vacancies/?category=Golang&exp=" + experience

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		var infoDou []models.Dou
		col.OnHTML("li.l-vacancy", func(e *colly.HTMLElement) {
			select {
			case <-ctx.Done():
				return
			default:
				position := e.ChildText("a.vt")
				href := e.ChildAttr("a.vt", "href")
				date := e.ChildText("div.date")
				info := models.Dou{
					Position: position,
					Href:     href,
					Date:     date,
				}
				infoDou = append(infoDou, info)
			}
		})

		col.OnError(func(r *colly.Response, err error) {
			logger.Errorf("Request URL: %s failed with response: %v\nError: %v\n", r.Request.URL, r, err)
		})

		err := col.Visit(URLDOU)
		if err != nil {
			return nil, err
		}
		chDou := make(chan []models.Dou, 1)
		chDou <- infoDou
		return chDou, nil
	}

}
