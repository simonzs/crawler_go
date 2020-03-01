package engine

import (
	"crawler_go/fetcher"
	"log"
)

// SimpleEngine 单任务引擎
type SimpleEngine struct{}

// Run 实现的是单任务爬虫架构
func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("Fetching %s", r.URL)

		parserResult, err := worker(r)
		if err != nil {
			continue
		}

		itemCount := 1

		requests = append(requests,
			parserResult.Reuqests...)

		for _, item := range parserResult.Items {
			log.Printf("Got item #%d: %v", itemCount, item)
			itemCount++
		}
	}
}

func worker(r Request) (ParserResult, error) {
	body, err := fetcher.Fetch(r.URL)
	if err != nil {
		log.Printf("Fetcher: error"+
			"fetching url %s: %v", r.URL, err)
		return ParserResult{}, err
	}

	return r.ParserFunc(body), nil
}
