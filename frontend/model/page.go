package model

// SearchResult ...
type SearchResult struct {
	Hits  int64
	Start int
	Items []interface{}
}
