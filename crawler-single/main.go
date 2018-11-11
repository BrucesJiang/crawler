package main

import (
	"crawler/crawler-single/zhenai/parser"
	"crawler/crawler-single/engine"
)


func main() {
	engine.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
