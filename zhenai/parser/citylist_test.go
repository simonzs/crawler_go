package parser

import (
	"io/ioutil"
	"testing"
)

const checkURL = "https://www.zhenai.com/zhenghun"

func TestParserCityList(t *testing.T) {

	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParserCityList(contents)

	const resultSize = 470
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	if len(result.Reuqests) != resultSize {
		t.Errorf("result should have %d"+
			"request; but had %d",
			resultSize, len(result.Reuqests))
	}
	for i, url := range expectedUrls {
		if result.Reuqests[i].URL != url {
			t.Errorf("excepted url #%d: %s, but "+
				"was %s", i, url, result.Reuqests[i].URL)
		}
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d"+
			"request; but had %d",
			resultSize, len(result.Items))
	}
}
