package main

import (
	"go-learning-app/pkg/web"
	"log"
	"fmt"
)

func main() {
	r, err := web.GetResponse("https://thenextweb.com/")
	defer r.Body.Close()

	if err != nil {
		log.Fatalln(err)
	}

	for e, v := range r.Header {
		fmt.Println(e + " - " + v[0])
	}

	web.Translate(r)
}
