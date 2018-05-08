package web

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetResponse(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		// handle error
		fmt.Println(err)

		return "N/A"
	}

	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)

	return string(contents)
}
