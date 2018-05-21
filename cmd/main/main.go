package main

import (
	"os"
	"fmt"
	"github.com/james-millner/go-lang-web-app/pkg/web"
)

func main() {

	fmt.Println("Starting Scrape...")
	fmt.Println(len(os.Args))

	url := os.Args[1]

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

	fmt.Println(len(links))

}