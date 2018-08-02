package web

import (
	"log"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/araddon/dateparse"
)

//RetreiveLinksFromDocument method
func RetreiveLinksFromDocument(url string, doc *goquery.Document) []string {
	var links []string

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, exists := s.Attr("href")

		//log.Printf("%s - %v", link, exists)
		if exists {
			//Only get links containing a protocol
			if strings.Contains(link, "http") && !ArrayContains(links, link) {
				links = append(links, link)
			}

			if !strings.Contains(link, "http") {
				formattedURL := url + link

				if strings.Contains(formattedURL, "http") && !ArrayContains(links, formattedURL) {
					links = append(links, formattedURL)
				}
			}

		}
	})

	return links
}

func GetFileName(file string) string {
	fileName, err := url.QueryUnescape(file)

	if err == nil {
		fileName = path.Base(fileName)
		fileName = strings.Replace(fileName, ".pdf", "", 1)
	}

	fileName = strings.Replace(fileName, "-", " ", -1)
	fileNames := strings.Split(fileName, "?")
	fileName = strings.Title(fileNames[0])
	return fileName
}

//IsPossibleCaseStudyLink function to determine if a given link mentions, or might be a case studies page.
func IsPossibleCaseStudyLink(url string) bool {

	casestudylinks := []string{"case-studies", "customers", "case_study", "stories", "case-study", "client-stories", "client-story", "wp-content"}

	for _, i := range casestudylinks {
		if strings.Contains(strings.ToLower(url), i) && isProbableLink(url) {
			return true
		}
	}

	return false
}

//IsPDFDocument function to find PDF links. This may be expanded on over time, hence the introduction of its own function.
func IsPDFDocument(url string) bool {
	return strings.Contains(url, ".pdf")
}

//IsProbableLink method
func IsProbableLink(url string) bool {
	return isProbableLink(url)
}

func isProbableLink(url string) bool {
	notInterestedIn := []string{"twitter", "https://t.co/", "youtube.com", "facebook.com", "linkedin.com", "mailto:", "terms-and-conditions", "T&C", "terms", "conditions",
		"privacy", "policy", "careers", "data-transfers", "pbs.twimg.com", "plus.google.com", "why", "login", "blog", "about-us", "myservices", "yahoo", "renewal", "campaign", "applications", "gartner", "features",
		"release", "consultancy", "google.com", "apple.com", "soundcloud.com", "news", "insights"}

	for _, p := range notInterestedIn {
		if strings.Contains(strings.ToLower(url), p) {
			return false
		}
	}

	return true
}

func TranslateMetaDataRowTime(metaRow string) time.Time {
	strs := strings.Split(metaRow, ",")
	str := strs[1]
	str = strings.Replace(str, "\"", "", -1)

	result, err := dateparse.ParseAny(str)

	if err != nil {
		log.Fatal(err)
		return time.Now()
	}

	return result
}
