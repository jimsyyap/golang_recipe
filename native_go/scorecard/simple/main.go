package main

import (
	"encoding/json"
	"github.com/google/go-github/v38/github"
	"golang.org/x/net/context"
	"log"
	"net/http"
)

func main() {
	// Create GitHub Client
	client := github.NewClient(&http.Client{})
	// Set Up Context
	ctx := context.Background()
	// Fetch Repository Information
	repo, _, err := client.Repositories.Get(ctx, "golang", "go")
	// Error Handling
	if err != nil {
		log.Fatalf(err.Error())
	}
	// **Convert to JSON**
	r, err := json.MarshalIndent(repo, "", "  ")
	// Error Handling for JSON Conversion
	if err != nil {
		log.Fatalf(err.Error())
	}
	// Print JSON
	log.Println(string(r))
}

/*
* this go code fetches details about the go repository owned by go on github and prints these details in a nicely formatted json format.
* the details include information like the repository's name, owner, description, stars, forks, issues, and more.
 */
