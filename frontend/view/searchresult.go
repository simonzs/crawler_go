package view

import (
	"crawler_go/frontend/model"
	"io"
	"text/template"
)

// SearchResultView 搜索视图
type SearchResultView struct {
	template *template.Template
}

// CreateSearchResultView 搜索视图
func CreateSearchResultView(
	filename string) SearchResultView {
	return SearchResultView{
		template: template.Must(template.ParseFiles(filename)),
	}
}

// Render 渲染
func (s SearchResultView) Render(
	w io.Writer, data model.SearchResult) error {
	return s.template.Execute(w, data)
}
