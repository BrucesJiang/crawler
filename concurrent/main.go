package main

import (
	"crawler/concurrent/engine"
	"crawler/concurrent/persist"
	"crawler/concurrent/scheduler"
	"crawler/concurrent/zhenai/parser"
)


func main() {
	//配置文件
	itemChan, err := persist.ItemSaver("dating_profile")

	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WokerCount:100,
		ItemChan: itemChan,
	}

	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}