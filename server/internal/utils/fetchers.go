package utils

import (
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gocolly/colly/v2"
)

func GenericFetch(method string, url string, body io.Reader) *http.Response {
	client := http.Client{
		Timeout: time.Second * 20,
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Fatal("Some oopsie occured:", err)
	}

	req.Header.Set("User-Agent", DefaultUserAgent)

	res, resErr := client.Do(req)
	if resErr != nil {
		log.Fatal("Some oopsie occured when getting a response:", resErr)
	}

	return res
}

type FetchHTMLOptions struct {
	Async bool
}

// A wrapper for `go-colly`
func FetchHTML(url string, opts *FetchHTMLOptions) string {
	var inhaler *colly.Collector

	if opts != nil && opts.Async {
		inhaler = colly.NewCollector(
			colly.UserAgent(DefaultUserAgent),
			colly.Async(),
		)
	} else {
		inhaler = colly.NewCollector(
			colly.UserAgent(DefaultUserAgent),
		)
	}

	var htmlContent string

	inhaler.OnHTML("html", func(e *colly.HTMLElement) {
		var err error
		htmlContent, err = e.DOM.Html()

		if err != nil {
			log.Fatal("Error extracting HTML content:", err)
		}
	})

	err := inhaler.Visit(url)
	if err != nil {
		log.Fatal("Error visiting URL:", err)
	}

	return htmlContent
}
