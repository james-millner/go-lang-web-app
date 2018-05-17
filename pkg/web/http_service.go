package web

import (
	"fmt"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)

type Content struct {
	Tag string
	Content string
}

func GetResponse(url string) (*goquery.Document, error)  {
	resp, err := http.Get(url)

	for e, v := range resp.Header {
		fmt.Println(e + " - " + v[0])
	}

	if err != nil {
		fmt.Errorf("failed to execute request: %v", err)
		return nil, err
	}

	return getDocument(resp)
}

func GetContent(document *goquery.Document, selector string, tag string) []Content {
	return translate(document, selector, tag);
}

func getDocument(response *http.Response) (*goquery.Document, error) {
	document, err := goquery.NewDocumentFromReader(response.Body)

	if err != nil {
		fmt.Errorf("couldn't read document: %v", err)
		return nil, err
	}

	return document, nil
}

func translate(doc *goquery.Document, selector string, t string) []Content {

	var results []Content

	// Find the review items
	doc.Find(selector).Each(func(i int, s *goquery.Selection) {

		var c = Content {Tag: t, Content: s.Find(t).Text()}
		results = append(results, c)
	})

	return results
}

