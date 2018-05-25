package web

import (
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"strings"
)

func getPageResponse(url string) *goquery.Document {
	r, err := GetResponse(url)

	if err != nil {
		errStr := fmt.Errorf("Couldn't get a response from: "+url, err)
		fmt.Println(errStr)
	}

	return r
}

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
		if strings.Contains(link, "http") && !contains(links, link) {
			links = append(links, link)
		}
	})

	return links
}

func CheckLinkHasSuffix(link string, suffix string) bool {
	return strings.HasSuffix(link, suffix)
}
