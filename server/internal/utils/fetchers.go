package utils

import (
	"io"
	"log"
	"net/http"
	"time"
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
