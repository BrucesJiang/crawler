package main

import (
	"crawler/concurrent/engine"
	"crawler/concurrent/persist"
	"crawler/concurrent/scheduler"
	"crawler/concurrent/zhenai/parser"
)


func main() {
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WokerCount:100,
		ItemChan: persist.ItemSaver(),
	}

	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}