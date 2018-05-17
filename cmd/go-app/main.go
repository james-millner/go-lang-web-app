package main

import (
	"os"
	"fmt"
	"go-learning-app/pkg/web"
)

func main() {

	url := os.Args[1]

	fmt.Println(url)

	r, err := web.GetResponse(url)

	if err != nil {
		fmt.Errorf("Couldn't get a response from: " + url, err)
	}

	doc := web.GetContent(r, ".u-row li .story", "a")

	for _, d := range doc {
		s := fmt.Sprintf("Tag %s, Content %s", d.Tag, d.Content)
		fmt.Println(s)
	}
}
