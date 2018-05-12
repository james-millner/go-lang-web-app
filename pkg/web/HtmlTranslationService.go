package web

import (
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"log"
	"fmt"
)

func Translate(response *http.Response) bool {
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
		return false
	}

	// Find the review items
	doc.Find(".u-row li .story").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Find("a").Text()
		title := s.Find("h4").Text()
		fmt.Printf("Link: %d: %s - %s\n", i, band, title)

	})

	return true
}
