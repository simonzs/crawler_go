package parser

import (
	"crawler_go/engine"
	"crawler_go/model"
	"io/ioutil"
	"testing"
)

func TestParserProfile(t *testing.T) {
	bytes, err := ioutil.ReadFile(
		"profile_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParserProfile(bytes,
		"https://album.zhenai.com/u/1451450381",
		"只等你")

	if len(result.Items) != 1 {
		t.Errorf("Result should contain 1 "+
			"elecment. but was %v", result.Items)
	}

	actual := result.Items[0]

	expected := engine.Item{
		URL:  "https://album.zhenai.com/u/1451450381",
		Type: "zhenai",
		ID:   "1451450381",
		Payload: model.Profile{
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

	if actual != expected {
		t.Errorf("expected %v; but was %v",
			expected, actual)
	}

}
