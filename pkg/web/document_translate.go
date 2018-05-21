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
		if strings.Contains(link, "http") && !contains(links, link) {
			links = append(links, link)
		}
	})

	return links
}

func CheckLinkHasSuffix(link string, suffix string) bool {
	if strings.HasSuffix(link, suffix) {
		return true
	}

	return false
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func split(str string, sep string) []string {
	return strings.Split(str, sep);
}
