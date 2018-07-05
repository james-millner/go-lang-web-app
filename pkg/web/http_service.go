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

func GetResponse(url string) (*goquery.Document, error) {
	resp, err := http.Get(url)

	if resp == nil {
		return nil, err
	}

	defer resp.Body.Close()

	if err != nil {
		fmt.Errorf("failed to execute request: %v", err)
		return nil, err
	}

	return getDocument(resp)
}
