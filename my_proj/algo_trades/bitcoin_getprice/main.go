package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const apiURL = "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd"

// response structure for the coingecko api
type PriceResponse struct {
	Bitcoin struct {
		USD float64 `json:"usd"`
	} `json:"bitcoin"`
}

func main() {
	//make an http GET request
	resp, err := http.Get(apiURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	//checi kf the status code is ok
	if resp.StatusCode != http.StatusOK {
		log.Fatal(resp.StatusCode)
	}

	//decode the response
	var priceResponse PriceResponse
	if err := json.NewDecoder(resp.Body).Decode(&priceResponse); err != nil {
		log.Fatal(err)
	}

	//print the price
	fmt.Println(priceResponse.Bitcoin.USD)
}
