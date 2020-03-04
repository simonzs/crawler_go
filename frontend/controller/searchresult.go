package controller

import (
	"context"
	"crawler_go/engine"
	"crawler_go/frontend/model"
	"crawler_go/frontend/view"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/olivere/elastic"
)

// TODO
// fill in query string
// support search button
// support paging
// add start page

// SearchResultHandler ...
type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

// CreateSearchResultHandler ...
func CreateSearchResultHandler(template string) SearchResultHandler {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return SearchResultHandler{
		view: view.CreateSearchResultView(
			template),
		client: client,
	}
}

// ServerHTTP localhost:9300/search?q=男 已购房 & from=20
func (h SearchResultHandler) ServeHTTP(
	w http.ResponseWriter, req *http.Request) {

	q := strings.TrimSpace(req.FormValue("q"))
	from, err := strconv.Atoi(req.FormValue("from"))
	if err != nil {
		from = 0
	}
	// fmt.Fprintf(w, "q=%s, from=%d", q, from)

	page, err := h.getSearchResult(q, from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = h.view.Render(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

// getSearchResult ...
func (h SearchResultHandler) getSearchResult(
	q string, from int) (model.SearchResult, error) {
	var result model.SearchResult
	resp, err := h.client.
		Search("dating_profile").
		Query(elastic.NewQueryStringQuery(
			rewriteQueryString(q))).
		From(from).
		Do(context.Background())
	if err != nil {
		return result, err
	}

	// result.Hits = resp.TotalHits()
	result.Hits = 12
	result.Start = from

	result.Items = resp.Each(
		reflect.TypeOf(engine.Item{}))

	// for _, v := range resp.Each(
	// 	reflect.TypeOf(engine.Item{})) {
	// 	result.Items = append(result.Items, v.(engine.Item))
	// }
	return result, nil
}

// rewriteQueryString ...
func rewriteQueryString(q string) string {
	re := regexp.MustCompile(`([A-Z][a-z]*):`)
	return re.ReplaceAllString(q, "Payload.$1:")
}
