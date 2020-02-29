package fetcher

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"testing"
)

func checkURL(url string, t *testing.T) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("wrong status code: %d", resp.StatusCode)
	}
}

func checkURLClient(url string, t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("User-Agent",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.122 Safari/537.36")
	if err != nil {
		panic(err)
	}
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", s)
}

func TestFetcher(t *testing.T) {
	const url = "http://album.zhenai.com/u/1727435860"
	// checkURL(url, t)
	checkURLClient(url, t)
}
