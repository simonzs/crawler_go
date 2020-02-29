package parser

import (
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
	result := ParserProfile(bytes, "海花")
	if len(result.Items) != 1 {
		t.Errorf("Result should contain 1 "+
			"elecment. but was %v", result.Items)
	}

	profile := result.Items[0].(model.Profile)

	excepted := model.Profile{
		Name:       "海花",
		Gender:     "女士",
		Age:        36,
		Height:     168,
		Income:     "3-5千",
		Marriage:   "离异",
		Education:  "高中及以下",
		Occupation: "阿坝小金",
	}

	if profile != excepted {
		t.Errorf("excepted %v; but was %v",
			excepted, profile)
	}

}
