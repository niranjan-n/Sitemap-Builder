package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/niranjan-n/HTML-Parser/link"
)

func main() {
	urlFlag := flag.String("url", "https://golang.org/", "The url that you want to build a sitemap for")
	flag.Parse()
	fmt.Println(*urlFlag)
	resp, err := http.Get(*urlFlag)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	reqURL := resp.Request.URL
	baseURL := &url.URL{
		Scheme: reqURL.Scheme,
		Host:   reqURL.Host,
	}
	base := baseURL.String()
	links, _ := link.Parse(resp.Body)
	var hrefs []string
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			hrefs = append(hrefs, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			hrefs = append(hrefs, l.Href)
		}
	}
	for _, href := range hrefs {
		fmt.Println(href)
	}

}
