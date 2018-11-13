package engine

import (
	"crawler/concurrent/fetcher"
	"log"
)

type SimpleEngine struct {

}


func (s SimpleEngine)Run(seeds ...Request) {
	requests := []Request{}

	//将种子放入
	for _, r := range seeds {
		requests = append(requests, r)
	}

	//如果队列中还有元素则继续循环
	for len(requests) >= 1 {

		//取元素
		request := requests[0]
		requests = requests[1:]

		result, err  := Worker(request)

		if err != nil {
			continue
		}

		requests = append(requests, result.Requests...)

		for _, item := range result.Items {
			log.Printf("Got item %s\n", item)
		}
	}
}


//Worker
func Worker(r Request) (ParseResult, error){
	log.Printf("Got Url :" + r.Url)

	//获取内容
	content, err := fetcher.Fetch(r.Url)

	if err != nil {
		log.Printf("Fetcher error, fetching is %s, %s\n", r.Url, err)
		return ParseResult{}, err
	}

	return r.ParseFunc(content), nil
}