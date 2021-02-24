package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"log"

	"github.com/gocolly/colly"
)

type Restaurant struct {
	Name    string
	Photo   string
	Cuisine string
	// Location string
	Summary string
}

func main() {

	selector := "body > table:nth-child(7) > tbody > tr > td:nth-child(1) > table:nth-child(17) > tbody > tr > td > table > tbody" // grabs just the body

	c := colly.NewCollector(
		colly.AllowedDomains("zabihah.com", "www.zabihah.com"),
	)

	infoCollector := c.Clone()

	c.OnHTML(selector, func(e *colly.HTMLElement) {
		restaurantUrl := e.ChildText("div.titleBS > a")
		restaurantUrl = e.Request.AbsoluteURL(restaurantUrl)
		infoCollector.Visit(restaurantUrl)

		e.ForEach("tr > td:nth-child(3)", func(_ int, h *colly.HTMLElement) { // for loop, to get each individual restaurant
			link := h.ChildAttr("a", "href")
			fmt.Printf("Link found: -> %s\n", link)
		})
	})

	infoCollector.OnHTML("body > table:nth-child(7) > tbody > tr > td", func(e *colly.HTMLElement) {
		tmpRestaurant := Restaurant{}
		tmpRestaurant.Name = e.ChildText("body > table:nth-child(7) > tbody > tr > td:nth-child(1) > table:nth-child(7) > tbody > tr > td:nth-child(1)")
		tmpRestaurant.Photo = e.ChildAttr("body > table:nth-child(7) > tbody > tr > td:nth-child(1) > table:nth-child(9) > tbody > tr:nth-child(2) > td > table > tbody > tr > td:nth-child(1) > img", "src")
		tmpRestaurant.Cuisine = e.ChildText("#alertBox2 > div > b")

		tmpRestaurant.Summary = strings.TrimSpace(e.ChildText("#alertBox2"))

		js, err := json.MarshalIndent(tmpRestaurant, "", "    ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(js))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL.String())
	})

	infoCollector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting Restaurant URL: ", r.URL.String())
	})

	// uncomment below line if you enable Async mode
	// c.Wait()
	startUrl := fmt.Sprintf("https://www.zabihah.com/sub/United-States/California/Sacramento/QE8Gfd3yNC")
	c.Visit(startUrl)
}
