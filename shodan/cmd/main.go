package main

import (
	"fmt"
	"log"
	"os"

	"github.com/blackhat-go/bhg/ch-3/shodan/shodan"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("usage: main <searchterm>")
	}
	apiKey := os.Getenv("SHODAN_API_KEY")
	s := shodan.New(apiKey)
	info, err := s.APIInfo()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Printf(
		"query credits remaining: %d\n",
		info.QueryCredits)
	// info.ScanCredits)

	hostSearch, err := s.HostSearch(os.Args[1])
	if err != nil {
		log.Panicln(err)
	}
	for _, host := range hostSearch.Matches {
		fmt.Println(host.IPString, host.Port)
	}
}

/*
* Key points:
* shodan is a search engine. Instead of websites, it finds computers and devices (IoT)
* search terms can be specific, like webcams, or general (servers)
* this code only finds computers; it doesnt interact with them
 */
