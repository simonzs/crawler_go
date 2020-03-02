package engine
import (
	"crawler_go/fetcher"
	"log"
)

func worker(r Request) (ParserResult, error) {
	body, err := fetcher.Fetch(r.URL)
	if err != nil {
		log.Printf("Fetcher: error"+
			"fetching url %s: %v", r.URL, err)
		return ParserResult{}, err
	}

	return r.ParserFunc(body, r.URL), nil
}