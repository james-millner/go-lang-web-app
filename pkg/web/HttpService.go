package web

import (
	"fmt"
	"net/http"
)

func GetResponse(url string) (*http.Response, error)  {
	resp, err := http.Get(url)
	if err != nil {
		// handle error
		fmt.Println(err)

		fmt.Errorf("failed to execute request: %v", err)
		return nil, err
	}

	return resp, nil
}
