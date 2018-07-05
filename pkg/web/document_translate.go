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

func GetLinks(doc *goquery.Document) []string {
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

func isValidCaseStudyLink(url string, expectingDocument bool) bool {
	caseStudyLink := false

	if expectingDocument {

		//Only want PDF links.
		if strings.Contains(url, ".pdf") {
			return true
		} 

		return caseStudyLink

	} 

	if !strings.Contains(url, "http") {
		return caseStudyLink
	}

	casestudylinks := []string{"case-studies", "customers"}

	for _, i := range casestudylinks {
		if(strings.Contains(url, i)) {
			caseStudyLink = true
		}
	}


	return caseStudyLink
	
}
