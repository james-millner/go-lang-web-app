package main

import (
	"os"
	"fmt"
	"go-learning-app/pkg/web"
)

func main() {

	url := os.Args[1]
	selector := os.Args[2]
	tag := os.Args[3]

	fmt.Println(url)

	r, err := web.GetResponse(url)

	if err != nil {
		errStr := fmt.Errorf("Couldn't get a response from: " + url, err)
		fmt.Println(errStr)
	}

	doc := web.GetContent(r, selector, tag)

	for _, d := range doc {
		s := fmt.Sprintf("Tag %s, Content %s", d.Tag, d.Content)
		fmt.Println(s)
	}
}