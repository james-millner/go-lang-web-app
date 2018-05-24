package main

import (
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/james-millner/go-lang-web-app/pkg/web"
	"github.com/gorilla/mux"
	"strings"
)

type Response struct {
	Success   bool
	Links     []string
	Documents []string
}

func main() {

	port := ":4000"

	router := mux.NewRouter()
	router.HandleFunc("/gather-links", GatherLinks).Methods("POST")

	fmt.Println("Listening on: ", port)
	http.ListenAndServe(port, router)

}

func GatherLinks(w http.ResponseWriter, r *http.Request) {

	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)

	url := r.FormValue("url")

	if url != "" {

		var links []string
		var documents []string

		for _, t := range getLinks(url) {
			if strings.Contains(t, ".pdf") {
				documents = append(documents, t)
			} else {
				links = append(links, t)
			}
		}

		resp := &Response{Links: links, Success: true, Documents: documents}
		enc.Encode(resp)
	} else {
		resp := &Response{Success: false}
		enc.Encode(resp)
	}
}

func getLinks(url string) []string {

	fmt.Println(url)

	r, err := web.GetResponse(url)

	if err != nil {
		errStr := fmt.Errorf("Couldn't get a response from: "+url, err)
		fmt.Println(errStr)
	}

	var links = web.GetPageLinks(r)

	for l := range links {
		fmt.Println(links[l])
	}

	fmt.Println("Total links found for", url, ":", len(links))
	return links
}
