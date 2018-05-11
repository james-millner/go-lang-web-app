package main

import (
	"fmt"
	"go-learning-app/pkg/web"
	"log"
)

func main() {
	r, err := web.GetResponse("https://www.google.co.uk")
	defer r.Body.Close()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("*** Header ***")
	for k, v := range r.Header {
		fmt.Println("Header field " + string(k))

		for a := range v {
			fmt.Println(a)
		}

	}

	fmt.Println("*** Body ***")

	fmt.Println(r.Body)
}
