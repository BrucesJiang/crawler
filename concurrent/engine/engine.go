package engine

import (
	"crawler/single/fetcher"
	"log"
)

func Run(seeds ...Request) {
	requests := []Request{}

	//将种子放入
	for _, r := range seeds {
		requests = append(requests, r)
	}

	//如果队列中还有元素则继续循环
	for len(requests) < 1 {

		//取元素
		request := requests[0]
		requests = requests[1:]

		log.Printf("Got Url :" + request.Url)

		//获取内容
		content, err := fetcher.Fetch(request.Url)

		if err != nil {
			log.Printf("Fetcher error, fetching is %s, %s\n", request.Url, err)
			continue
		}

		parseResult := request.ParseFunc(content)
		//
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item %s\n", item)
		}
	}
}