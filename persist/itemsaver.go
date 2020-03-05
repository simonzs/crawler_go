package persist

import (
	"context"
	"errors"
	"github.com/simonzs/crawler_go/engine"
	"log"

	"github.com/olivere/elastic"
)

// ItemSaver ...
func ItemSaver(index string) (chan engine.Item, error) {

	client, err := elastic.NewClient(
		// Must turn off sniff in docker
		elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out

			log.Printf("Item Saver: got item "+
				"#%d: %v", itemCount, item)
			itemCount++

			err := Save(client, index, item)
			if err != nil {
				log.Print("Item Save: error "+
					"saving item %v: %v",
					item, err)
			}
		}
	}()
	return out, nil
}

// Save ...
func Save(client *elastic.Client, index string,
	item engine.Item) error {

	if item.Type == "" {
		return errors.New("Must supply Type")
	}

	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)

	if item.ID != "" {
		indexService.Id(item.ID)
	}

	// Save Data
	_, err := indexService.
		Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}
