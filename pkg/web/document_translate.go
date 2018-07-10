package web

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//RetreiveLinksFromDocument method
func RetreiveLinksFromDocument(doc *goquery.Document) []string {
	var links []string

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, exists := s.Attr("href")

		if exists {
			//Only get links containing a protocol
			if strings.Contains(link, "http") && !ArrayContains(links, link) {
				links = append(links, link)
			}
		}
	})

	return links
}

//IsPossibleCaseStudyLink function to determine if a given link mentions, or might be a case studies page.
func IsPossibleCaseStudyLink(url string) bool {

	caseStudyLink := false

	casestudylinks := []string{"case-studies", "customers", "case_study", "stories", "case-study"}

	for _, i := range casestudylinks {
		if strings.Contains(strings.ToLower(url), i) {
			caseStudyLink = true
			break
		}
	}

	return caseStudyLink
}

//IsPDFDocument function to find PDF links. This may be expanded on over time, hence the introduction of its own function.
func IsPDFDocument(url string) bool {
	return strings.HasSuffix(url, ".pdf")
}

//IsProbableLink method
func IsProbableLink(url string) bool {

	notInterestedIn := []string{"twitter", "https://t.co/", "youtube.com", "facebook.com", "linkedin.com", "mailto:", "terms-and-conditions", "T&C", "terms", "conditions", "privacy", "policy", "careers", "data-transfers", "pbs.twimg.com", "plus.google.com", "solutions", "why", "login", "blog", "about-us"}

	for _, p := range notInterestedIn {
		if strings.Contains(strings.ToLower(url), p) {
			return false
		}
	}

	return true
}
