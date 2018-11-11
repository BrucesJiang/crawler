package main

import (
	"crawler/concurrent/engine"
	"crawler/concurrent/zhenai/parser"
)


func main() {
	engine.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}