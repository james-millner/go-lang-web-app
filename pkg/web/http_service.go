package web

import (
	"fmt"
	"io"

	"github.com/PuerkitoBio/goquery"
)

//GetGoqueryDocument method
func GetGoqueryDocument(r io.Reader) (*goquery.Document, error) {
	d, err := goquery.NewDocumentFromReader(r)

	if err != nil {
		e := fmt.Errorf("couldn't read document: %v", err)
		fmt.Println(e)
		return nil, err
	}

	return d, nil
}
