package persist

import (
	"context"
	"crawler_go/engine"
	"errors"
	"log"

	"github.com/olivere/elastic"
)

// ItemSaver ...
func ItemSaver() chan engine.Item {
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out

			log.Printf("Item Saver: got item "+
				"#%d: %v", itemCount, item)
			itemCount++

			err := save(item)
			if err != nil {
				log.Print("Item Save: error "+
					"saving item %v: %v",
					item, err)
			}
		}
	}()
	return out
}

func save(item engine.Item) error {
	client, err := elastic.NewClient(
		// Must turn off sniff in docker
		elastic.SetSniff(false))
	if err != nil {
		return err
	}

	if item.Type == "" {
		return errors.New("Must supply Type")
	}

	indexService := client.Index().
		Index("dating_profile").
		Type(item.Type).
		BodyJson(item)

	if item.ID != "" {
		indexService.Id(item.ID)
	}

	// Save Data
	_, err = indexService.
		Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}
