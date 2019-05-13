package persist

import (
	"context"
	"crawler/engine"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			itemCount++
			err := Save(client, index, item)
			if err != nil {
				log.Print("Item Saver: error, saving item %v:%v",
					item, err)
			}
		}
	}()

	return out, nil
}

func Save(client *elastic.Client,
	index string, item engine.Item) error {

	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err := indexService.
		Do(context.Background())

	return err
}
