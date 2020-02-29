package engine

import (
	"crawler_go/fetcher"
	"log"
)

// Run ...
func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("Fetching %s", r.URL)

		body, err := fetcher.Fetch(r.URL)
		if err != nil {
			log.Printf("Fetcher: error"+
				"fetching url %s: %v", r.URL, err)
			continue
		}
		parserResult := r.ParserFunc(body)
		requests = append(requests,
			parserResult.Reuqests...)

		for _, item := range parserResult.Items {
			log.Printf("Got item %s", item)
		}
	}
}
