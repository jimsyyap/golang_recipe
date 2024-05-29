package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Status struct {
	Message string
	Status  string
}

func main() {
	res, err := http.Post(
		"http://127.0.0.1:3790/import",
		"application/json",
		nil,
	)
	if err != nil {
		log.Fatalln(err)
	}

	var status Status
	if err := json.NewDecoder(res.Body).Decode(&status); err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	log.Printf("%s -> %s\n", status.Status, status.Message)
}

/*
* this program sends a post request to a specified URL, expects a JSON response containing 'status' and 'message' fields, decodes this response into a 'status' struct, and logs these fields.
* it's a simple example of making http requests and handling json responses in go
 */
