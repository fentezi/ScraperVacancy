package pkg

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

type InfoDjinni struct {
	Position string `json:"position"`
	Href     string `json:"href"`
	Date     string `json:"date_published"`
	Views    string `json:"views"`
	Reviews  string `json:"reviews"`
}

func ParseDjinni() []InfoDjinni {
	col := colly.NewCollector()
	var infoDjinni []InfoDjinni

	const baseUrl = "https://djinni.co/jobs/?primary_keyword=Golang"

	col.OnHTML("div.job-list-item", func(e *colly.HTMLElement) {
		position := strings.TrimSpace(e.ChildText("a.h3.job-list-item__link"))
		if strings.HasPrefix(position, "Junior") {
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
	})

	currentPage := 1
	col.OnScraped(func(response *colly.Response) {
		fmt.Println("Finished scraping:", response.Request.URL)
		if currentPage < 4 {
			currentPage++
		}
		nextUrl := fmt.Sprintf("%s&page=%d", baseUrl, currentPage)
		col.Visit(nextUrl)
	})

	col.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	col.Visit(fmt.Sprintf("%s&page=1", baseUrl))
	return infoDjinni
}
