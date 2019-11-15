package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Request /hello over port 8080 via the GET method
	// r, err := http.Get("http://localhost:8080/hello")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	
	// Request /hello over HTTPS port 8443 via the GET method
	r, err := http.Get("https://localhost:8443/hello")
	if err != nil {
		log.Fatal(err)
	}

	// Read the response body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Print the response body to stdout
	fmt.Printf("%s\n", body)
}