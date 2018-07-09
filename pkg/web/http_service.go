package web

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

//GetResponse method for retuning a GoQuery document for analysis.
func GetResponse(url string) (*goquery.Document, error) {
	resp, err := http.Get(url)

	if resp == nil {
		return nil, err
	}

	defer resp.Body.Close()

	if err != nil {
		e := fmt.Errorf("failed to execute request: %v", err)
		fmt.Println(e)
		return nil, err
	}

	document, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		e := fmt.Errorf("couldn't read document: %v", err)
		fmt.Println(e)
		return nil, err
	}

	return document, nil
}
