package persist

import (
	"context"
	"crawler_go/engine"
	"crawler_go/model"
	"encoding/json"
	"testing"

	"github.com/olivere/elastic"
)

func TestSave(t *testing.T) {
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

	// TODO: Try to start up elastic search
	// here using docker go client.
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	const index = "dating_test"
	// Save excepted item
	err = save(client, index, expected)
	if err != nil {
		panic(err)
	}

	// Fetch saved item
	resp, err := client.Get().
		Index(index).
		Type(expected.Type).
		Id(expected.ID).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	t.Logf("%s", resp.Source)

	var actual engine.Item

	bytes, _ := resp.Source.MarshalJSON()
	err = json.Unmarshal(bytes, &actual)

	if err != nil {
		panic(err)
	}

	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	// expected.Payload.Age = 10
	// Verify Result
	if expected != actual {
		t.Errorf("got %v; expected %v",
			actual, expected)
	}
}
