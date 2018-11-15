package persist

import (
	"context"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})

	go func() {
		itemCount := 0
		for {
			item := <- out
			log.Printf("Item Saver #%d, %v", itemCount, item)
			itemCount++
			save(item)
		}
	}()
	return out
}

func save(item interface{}) {
	client, err := elastic.NewClient(
		//must turn off in Docker
		elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	resp, err := client.Index().
		Index("dating_profile").
		Type("zhenai").
		BodyJson(item).
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	log.Printf("%+v", resp)
}