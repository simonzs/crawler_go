package view

import (
	"crawler_go/engine"
	"crawler_go/frontend/model"
	common "crawler_go/model"
	"os"
	"testing"
)

func TestSearchResultView_Render(t *testing.T) {
	// template := template.Must(
	// 	template.ParseFiles("template.html"))
	view := CreateSearchResultView(
		"template.html")

	out, err := os.Create("template.test.html")
	if err != nil {
		panic(out)
	}
	page := model.SearchResult{}
	page.Hits = 123
	page.Start = 1

	item := engine.Item{
		URL:  "https://album.zhenai.com/u/1451450381",
		Type: "zhenai",
		ID:   "1451450381",
		Payload: common.Profile{
			Name:       "只等你",
			Gender:     "女士",
			Age:        27,
			Height:     160,
			Income:     "8千-1.2万",
			Marriage:   "未婚",
			Education:  "大学本科",
			Occupation: "成都双流区",
		},
	}
	for i := 0; i < 100; i++ {
		page.Items = append(page.Items, item)
	}

	err = view.Render(out, page)
	// err = template.Execute(out, page)
	if err != nil {
		panic(err)
	}
}
