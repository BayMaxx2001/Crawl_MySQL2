package utils

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func RequestHTTP(url string) (*http.Response, error) {

	// Request http
	log.Println("Request http get to", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error at RequestHTTP of utils/response_HTTP: ", err)
		return nil, err
	}
	return resp, err
}

func LoadHTML(resp *http.Response) (*goquery.Document, error) {
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Println("Error at LoadHTML of utils/response_HTTP: ", err)
		return nil, err
	}
	return doc, err
}
