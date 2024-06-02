package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/jimsyyap/bingmetadata/metadata"
)

/* this function will process each search result, download the file and extracting metadata */
func handler(i int, s *goquery.Selection) {
	url, ok := s.Find("a").Attr("href")
	if !ok {
		return
	}

	fmt.Printf("%d: %s\n", i, url)
	res, err := http.Get(url)
	if err != nil {
		return
	}

	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}
	defer res.Body.Close()

	r, err := zip.NewReader(bytes.NewReader(buf), int64(len(buf)))
	if err != nil {
		return
	}

	cp, ap, err := metadata.NewProperties(r)
	if err != nil {
		return
	}

	log.Printf(
		"%21s %s - %s %s\n",
		cp.Creator,
		cp.LastModifiedBy,
		ap.Application,
		ap.GetMajorVersion())
}

/* sets up the search query, sends the req to bing and process results */
func main() {
	if len(os.Args) != 3 {
		log.Fatalln("Missing required argument. Usage: main.go <domain> <ext>")
	}
	domain := os.Args[1]
	filetype := os.Args[2]

	q := fmt.Sprintf(
		"site:%s && filetype:%s && instreamset:(url title):%s",
		domain,
		filetype,
		filetype)

	search := fmt.Sprintf("http://www.bing.com/search?q=%s", url.QueryEscape(q))
	res, err := http.Get(search)
	if err != nil {
		return
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Panicln(err)
	}
	defer res.Body.Close()
	s := "html body div#b_content ol#b_results li.b_algo h2"
	doc.Find(s).Each(handler)
}

/*
* imagine you want to find and download specific types of files (eg, pdf's) from a particular website. This program can do the following:
* search bing for those files on the specified website
* download the files
* reads the file and extracns information about those who created it, who modified and what application was used to create it
* prints out this information
* this program automates the process of searching, download and reading metadataa from files found on the internet
 */
