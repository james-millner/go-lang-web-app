package web

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
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

func RetreiveLinksFromDocument(doc *goquery.Document) []string {
	var links []string

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")

		//Only get links containing a protocol
		if strings.Contains(link, "http") && !ArrayContains(links, link) {
			links = append(links, link)
		}
	})

	return links
}

//IsPossibleCaseStudyLink function to determine if a given link mentions, or might be a case studies page.
func IsPossibleCaseStudyLink(url string) bool {

	caseStudyLink := false

	casestudylinks := []string{"case-studies", "customers"}

	for _, i := range casestudylinks {
		if strings.Contains(url, i) {
			caseStudyLink = true
		}
	}

	return caseStudyLink
}

//IsPDFDocument function to find PDF links. This may be expanded on over time, hence the introduction of its own function.
func IsPDFDocument(url string) bool {
	return strings.HasSuffix(url, ".pdf")
}

func IsProbableLink(url string) bool {

	notInterestedIn := []string{"twitter", "https://t.co/", "youtube.com", "facebook.com", "linkedin.com", "mailto:", "terms-and-conditions", "T&C", "terms", "conditions", "privacy", "policy", "careers", "data-transfers", "pbs.twimg.com", "plus.google.com"}

	for _, p := range notInterestedIn {
		if strings.Contains(url, p) {
			return false
		}
	}

	return true;
}