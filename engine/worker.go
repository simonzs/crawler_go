package engine

import (
	"log"

	"github.com/simonzs/crawler_go/fetcher"
)

// Worker ...
func Worker(r Request) (ParserResult, error) {
	body, err := fetcher.Fetch(r.URL)
	if err != nil {
		log.Printf("Fetcher: error"+
			"fetching url %s: %v", r.URL, err)
		return ParserResult{}, err
	}
	return r.Parser.Parse(body, r.URL), nil
}
