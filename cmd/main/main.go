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

type IndividualLinkResponse struct {
	Url			string
	Selector	string
	Tag			string
}

func main() {

	port := ":4000"

	router := mux.NewRouter()
	router.HandleFunc("/gather-links", gatherLinks).Methods("POST")
	router.HandleFunc("/handle-link", handle).Methods("POST")

	fmt.Println("Listening on: ", port)
	http.ListenAndServe(port, router)

}

func gatherLinks(w http.ResponseWriter, r *http.Request) {

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

		for l := range links {
			fmt.Println(links[l])
		}
	
		fmt.Println("Total links found for:", len(links))
		fmt.Println("Total documents found for:", len(documents))


		enc.Encode(resp)

	} else {

		resp := &Response{Success: false}
		enc.Encode(resp)

	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	selector := r.FormValue("selector")
	tag := r.FormValue("tag")

	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)

	resp := &IndividualLinkResponse{Url: url, Selector: selector, Tag: tag}
	enc.Encode(resp)
}


func getLinks(url string) []string {

	fmt.Println(url)

	r, err := web.GetResponse(url)

	if err != nil {

		error := fmt.Errorf("failed to execute request: %v", err)
		fmt.Println(error)

		return nil
	}

	return web.GetPageLinks(r)
}
