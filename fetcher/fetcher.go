package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

// Fetcher 提取器

var rateLimiter = time.Tick(100 * time.Millisecond)

// Fetch 提取文本
func Fetch(url string) ([]byte, error) {
	<-rateLimiter

	// resp, err := http.Get(url)

	log.Printf("Fetching %v", url)
	resp, err := GetByClient(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil,
			fmt.Errorf("wrong status code: %d",
				resp.StatusCode)
	}

	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

// GetByClient ...
func GetByClient(url string) (resp *http.Response, err error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("User-Agent",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.122 Safari/537.36")
	if err != nil {
		panic(err)
	}
	client := http.Client{}
	return client.Do(request)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error : %s", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
