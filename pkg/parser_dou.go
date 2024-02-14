package pkg

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

type InfoDou struct {
	Position string `json:"position"`
	Href     string `json:"href"`
	Date     string `json:"date_published"`
}

func ParserDou() []InfoDou {
	col := colly.NewCollector()
	inforDou := []InfoDou{}

	col.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36"

	const URLDOU = "https://jobs.dou.ua/vacancies/?category=Golang&exp=1-3"

	col.OnHTML("li.l-vacancy", func(e *colly.HTMLElement) {
		position := e.ChildText("a.vt")
		href := e.ChildAttr("a.vt", "href")
		date := e.ChildText("div.date")
		info := InfoDou{
			Position: position,
			Href:     href,
			Date:     date,
		}
		inforDou = append(inforDou, info)
	})

	col.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", string(r.Body), "\nError:", err)
	})

	col.Visit(URLDOU)
	return inforDou
}
