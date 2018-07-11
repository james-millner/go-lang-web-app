package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

//HandleLink function.
func (rs *ResponseService) HandleLink() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)

		url := r.FormValue("url")
		log.Println(url)

		fileName, _ := DownloadFile(url)

		f, err := os.Open(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		body, err := rs.rs.TikaClient.Parse(context.Background(), f)

		if err != nil {
			e := fmt.Errorf("Error with TikaClient parse: %v", err)
			log.Fatal(e)
		}

		os.Remove(fileName)
		log.Println(body)

	}
}

//DownloadFile function
func DownloadFile(url string) (string, error) {

	tokens := strings.Split(url, "/")
	fileName := tokens[len(tokens)-1]

	out, oserr := os.Create(fileName)

	if oserr != nil {
		e := fmt.Errorf("Error with creating OS file: %v", oserr)
		log.Fatal(e)
		return "", oserr
	}

	resp, err := http.Get(url)
	if err != nil {
		e := fmt.Errorf("Error with GET request: %v", err)
		log.Fatal(e)
		return "", oserr
	}
	defer resp.Body.Close()

	io.Copy(out, resp.Body)

	return fileName, nil
}
