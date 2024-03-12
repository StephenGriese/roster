package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

type Player struct {
	url, image, name, price string
}

func main() {

	// creating a new Colly instance
	c := colly.NewCollector()

	// scraping logic
	c.OnHTML("div", func(e *colly.HTMLElement) {
		fmt.Printf("e: %v", e)
	})

	// visiting the target page
	c.Visit("https://www.nhl.com/flyers/roster")

}
