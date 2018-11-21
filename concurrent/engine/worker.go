package engine

import (
	"crawler/concurrent/fetcher"
	"log"
)

//Worker
func Worker(r Request) (ParseResult, error) {
	log.Printf("Got Url :" + r.Url)

	//获取内容
	content, err := fetcher.Fetch(r.Url)

	if err != nil {
		log.Printf("Fetcher error, fetching is %s, %s\n", r.Url, err)
		return ParseResult{}, err
	}

	return r.ParseFunc(content), nil
}
