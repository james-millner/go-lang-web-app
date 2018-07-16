package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

//ProcessLink function.
func (rs *ResponseService) ProcessLink() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)

		url := r.FormValue("url")
		log.Println(url)

		tokens := strings.Split(url, "/")
		fileName := tokens[len(tokens)-1]
	
		out, oserr := os.Create(fileName)
	
		if oserr != nil {
			e := fmt.Errorf("Error with creating OS file: %v", oserr)
			log.Fatal(e)
		}
	
		resp, err := http.Get(url)
		if err != nil {
			e := fmt.Errorf("Error with GET request: %v", err)
			log.Fatal(e)
		}
		
		defer resp.Body.Close()
	
		io.Copy(out, resp.Body)

		f, err := os.Open(fileName)
		if err != nil {
			log.Fatal(err)
		} else {
			body, err := rs.rs.TikaClient.Parse(context.Background(), f)

			if err != nil {
				e := fmt.Errorf("Error with TikaClient parse: %v", err)
				log.Fatal(e)
			} else {
	
				body := strings.TrimSpace(body)

				//bodyArray := strings.Split(body, "/n")
	
				f.Close()
				os.Remove(fileName)
				log.Println(body)
			}
		}
	}
}

