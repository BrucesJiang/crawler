package main

import (
	"crawler/concurrent/engine"
	"crawler/concurrent/zhenai/parser"
	"crawler/concurrent/scheduler"
)


func main() {
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
		WokerCount:10,
	}

	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}