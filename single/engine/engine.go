package engine

import (
	"crawler/single/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parseResult, err := Worker(r) //抽象出来Worker

		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("Got item %+v\n", item)
		}
	}
}

func Worker(r Request) (ParseResult, error) {

	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error " +
			"fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}
	parseResult := r.ParserFunc(body)
	return parseResult, nil
}