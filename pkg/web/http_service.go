package web

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Content struct {
	Tag     string
	Content string
}

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

	return getDocument(resp)
}
