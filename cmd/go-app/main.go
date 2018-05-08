package main

import (
	"fmt"
	"go-learning-app/pkg/web"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println(web.GetResponse("https://www.google.co.uk"))
}
