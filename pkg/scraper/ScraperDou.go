package scraper

import (
	"context"
	"github.com/fentezi/scraper/pkg/logging"
	"sync"

	"github.com/gocolly/colly/v2"
)

type InfoDou struct {
	Position string `json:"position"`
	Href     string `json:"href"`
	Date     string `json:"date_published"`
}

func ParserDou(ctx context.Context, logger logging.Logger, chDou chan interface{}, experience string, wg *sync.WaitGroup) {
	col := colly.NewCollector()
	infoDou := []InfoDou{}
	switch experience {
	case "Junior":
		experience = "0-1"
	case "Middle":
		experience = "1-3"
	case "Senior":
		experience = "3-5"
	}

	defer wg.Done()

	col.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36"

	URLDOU := "https://jobs.dou.ua/vacancies/?category=Golang&exp=" + experience

	col.OnHTML("li.l-vacancy", func(e *colly.HTMLElement) {
		select {
		case <-ctx.Done():
			return
		default:
			position := e.ChildText("a.vt")
			href := e.ChildAttr("a.vt", "href")
			date := e.ChildText("div.date")
			info := InfoDou{
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

	col.Visit(URLDOU)
	chDou <- infoDou
}
