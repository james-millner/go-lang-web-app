package web

import (
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"strings"
)

func getDocument(response *http.Response) (*goquery.Document, error) {
	document, err := goquery.NewDocumentFromReader(response.Body)

	if err != nil {
		fmt.Errorf("couldn't read document: %v", err)
		return nil, err
	}

	return document, nil
}

func getLinks(doc *goquery.Document) []string {
	var links []string

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")

		//Only get links containing a protocol
		if strings.Contains(link, "http") {
			links = append(links, link)
		}
	})

	return links
}
