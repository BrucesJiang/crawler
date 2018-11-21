package persist

import (
	"context"
	"crawler/concurrent/engine"
	"errors"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func ItemSaver(index string) (chan engine.Item, error ) {
	out := make(chan engine.Item)
	client, err := elastic.NewClient(
		//must turn off in Docker
		elastic.SetSniff(false))

	if err != nil {
		//panic(err)
		return nil, err
	}
	go func() {
		itemCount := 0
		for {
			item := <- out
			log.Printf("Item Saver #%d, %v", itemCount, item)
			itemCount++
			err = save(client, index, item)

			if err != nil {
				log.Printf("Item saver: error " + " saveing %v: %v", item, err)
			}
		}
	}()
	return out, nil
}

func save(
	client *elastic.Client, index string,
	item engine.Item)  error{
	if item.Type == "" {
		return errors.New("must supply Type ")
	}

	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err := indexService.Do(context.Background())
	if err != nil {
		return  err
	}
	//log.Printf("%+v", resp)
	return nil
}