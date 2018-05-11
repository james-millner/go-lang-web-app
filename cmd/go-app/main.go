package main

import (
	"go-learning-app/pkg/web"
	"log"
	"fmt"
	"io/ioutil"
)

func main() {
	r, err := web.GetResponse("https://www.twitter.com")
	defer r.Body.Close()

	if err != nil {
		log.Fatalln(err)
	}

	for e, v := range r.Header {
		fmt.Println(e + " - " + v[0])
	}

	fmt.Println("***")

	contents, err := ioutil.ReadAll(r.Body)

	println(string(contents))

}
