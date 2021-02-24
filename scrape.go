package main

import (
	"encoding/json"
	"fmt"

	"log"

	"github.com/gocolly/colly"
)

type restaurant struct {
	Name    string
	Cuisine string
	Rating  string
}

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {
	selector := "body > table:nth-child(7) > tbody > tr > td:nth-child(1) > table:nth-child(17) > tbody > tr > td > table > tbody" // grabs just the body

	// Instantiate default collector
	c := colly.NewCollector()

	// On every a element which has href attribute call callback
	c.OnHTML(selector, func(e *colly.HTMLElement) {

		tmpRestaurant := restaurant{}
		tmpRestaurant.Name = e.ChildText("div.titleBS > a")
		tmpRestaurant.Cuisine = e.ChildText("#alertBox2 > div")
		tmpRestaurant.Rating = e.ChildText("#badge_score")

		// link := e.Attr("href")
		e.ForEach("tr > td:nth-child(3)", func(_ int, h *colly.HTMLElement) { // for loop, to get each individual restaurant
			link := h.ChildAttr("a", "href")
			fmt.Printf("Link found: -> %s\n", link)
		})
		js, err := json.MarshalIndent(tmpRestaurant, "", "    ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(js))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	//Start scraping https://www.zabihah.com/sub/United-States/California/Sacramento/QE8Gfd3yNC
	c.Visit("https://www.zabihah.com/sub/United-States/California/Sacramento/QE8Gfd3yNC")

	// // Create output file
	// outputFile, err := os.Create("output.html")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// //Copy data from HTTP response to file
	// _, err = io.Copy(outputFile, link.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
