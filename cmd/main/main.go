package main

import (
		"fmt"

	"github.com/james-millner/go-lang-web-app/pkg/web"
	"github.com/gorilla/mux"
		"net/http"
	"encoding/json"
)

type Response struct {
	Success bool
	Results []string
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
		links := getLinks(url)

		resp := &Response{Results: links, Success: true}
		enc.Encode(resp)
	} else {
		resp := &Response{ Success: false}
		enc.Encode(resp)
	}
}

func getLinks(url string) []string {

	fmt.Println(url)

	r, err := web.GetResponse(url)

	if err != nil {
		errStr := fmt.Errorf("Couldn't get a response from: " + url, err)
		fmt.Println(errStr)
	}

	var links = web.GetPageLinks(r)

	for l := range links {
		fmt.Println(links[l])
	}

	fmt.Println("Total links found for", url, ":", len(links))
	return links
}