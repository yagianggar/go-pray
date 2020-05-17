package pkpu

import (
	"fmt"
	"github.com/gocolly/colly"
)

type PKPUService struct {
	URL string
}

func (ps *PKPUService) GetPrayingSchedule() []Schedule {
	return parseHTML(ps.URL)
}

type Schedule struct {
	Date string
	Shubuh string
	Dzuhur string
	Ashar string
	Maghrib string
	Isya string
}

func parseHTML(url string) []Schedule {
	c := colly.NewCollector()
	schedules := []Schedule{}

	c.OnHTML("tbody", func(e *colly.HTMLElement)  {
		e.ForEach("tr", func(i int, trEl *colly.HTMLElement) {
			trClass := trEl.Attr("class")
			if trClass == "table_light" || trClass == "table_dark" || trClass == "table_highlight" {
				s := Schedule{}
				trEl.ForEach("td", func(j int, tdEl *colly.HTMLElement) {
					if j == 0 {
						s.Date = tdEl.Text
					} else if j == 1 {
						s.Shubuh = tdEl.Text
					} else if j == 2 {
						s.Dzuhur = tdEl.Text
					} else if j == 3 {
						s.Ashar = tdEl.Text
					} else if j == 4 {
						s.Maghrib = tdEl.Text
					} else if j == 5 {
						s.Isya = tdEl.Text
					}
				})
				schedules = append(schedules, s)
			}
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.Visit(url)
	return schedules
}

func NewService(url string) PKPUService {
	return PKPUService{
		URL: url,
	}
}