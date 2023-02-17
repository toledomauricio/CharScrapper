package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	var characters []string

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		name := e.Text

		if len(name) > 0 && name[0] >= 'A' && name[0] <= 'Z' {
			characters = append(characters, name)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("")

	for _, name := range characters {
		fmt.Println(name)
	}
}
